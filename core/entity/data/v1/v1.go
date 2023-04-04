/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                    <aegis.ist>
 *     .\_/.
 */

package v1

import (
	"encoding/json"
	"fmt"
	"github.com/shieldworks/aegis/core/template"
	"time"
)

type JsonTime time.Time

func (t JsonTime) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format(time.RubyDate))
	return []byte(stamp), nil
}

type Secret struct {
	Name    string   `json:"name"`
	Created JsonTime `json:"created"`
	Updated JsonTime `json:"updated"`
}

type BackingStore string

var File BackingStore = "file"
var Memory BackingStore = "memory"

type SecretFormat string

var Json SecretFormat = "json"
var Yaml SecretFormat = "yaml"
var None SecretFormat = "none"

type SecretMeta struct {
	// Overrides Env.SafeUseKubernetesSecrets()
	UseKubernetesSecret bool `json:"k8s"`
	// Overrides Env.SafeBackingStoreType()
	BackingStore BackingStore `json:"storage"`
	// Defaults to "default"
	Namespace string `json:"namespace"`
	// Go template used to transform the secret.
	// Sample secret:
	// '{"username":"admin","password":"AegisRocks"}'
	// Sample template:
	// '{"USER":"{{.username}}", "PASS":"{{.password}}"}"
	Template string `json:"template"`
	// Defaults to None
	Format SecretFormat
	// For tracking purposes
	CorrelationId string `json:"correlationId"`
}

type SecretStored struct {
	// Name of the secret.
	Name string
	// Raw value.
	Value string
	// Transformed value. This value is the value that workloads see.
	//
	// Apply transformation (if needed) and then store the value in
	// one of the supported formats. If the format is json, ensure that
	// a valid JSON is stored here. If the format is yaml, ensure that
	// a valid YAML is stored here. If the format is none, then just
	// apply transformation (if needed) and do not do any validity check.
	ValueTransformed string `json:"valueTransformed"`
	// Additional information that helps formatting and storing the secret.
	Meta SecretMeta
	// Timestamps
	Created time.Time
	Updated time.Time
}

// ToMapForK8s returns a map that can be used to create a Kubernetes secret.
//
//  1. If there is no template, ttempt to unmarshal the secret’ss value
//     into a map. If that fails, store the secret’s value under the "VALUE" key.
//  2. If there is a template, attempt to parse it. If parsing is successful,
//     create a new map with the parsed data. If parsing fails, follow the same
//     logic as in case 1, attempting to unmarshal the secret’s value into a map,
//     and if that fails, storing the secret’s value under the "VALUE" key.
func (secret SecretStored) ToMapForK8s() map[string][]byte {
	data := make(map[string][]byte)

	// If there is no template, use the secret’s value as is.
	if secret.Meta.Template == "" {
		err := json.Unmarshal(([]byte)(secret.Value), &data)
		if err != nil {
			value := secret.Value
			data["VALUE"] = ([]byte)(value)
		}

		return data
	}

	// Otherwise, apply the template.
	newData, err := template.ParseForK8sSecret(secret)
	if err == nil {
		data = make(map[string][]byte)
		for k, v := range newData {
			data[k] = ([]byte)(v)
		}

		return data
	}

	// If the template fails, use the secret’s value as is.
	err = json.Unmarshal(([]byte)(secret.Value), &data)
	if err != nil {
		value := secret.Value
		data["VALUE"] = ([]byte)(value)
	}

	return data
}

// ToMap converts the SecretStored struct to a map[string]any.
// The resulting map contains the following key-value pairs:
//
//	"Name": the Name field of the SecretStored struct
//	"Value": the Value field of the SecretStored struct
//	"Created": the Created field of the SecretStored struct
//	"Updated": the Updated field of the SecretStored struct
func (secret SecretStored) ToMap() map[string]any {
	return map[string]any{
		"Name":    secret.Name,
		"Value":   secret.Value,
		"Created": secret.Created,
		"Updated": secret.Updated,
	}
}
