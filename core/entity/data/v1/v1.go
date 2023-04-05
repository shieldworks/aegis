/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                    <aegis.ist>
 *     .\_/.
 */

package v1

import (
	"bytes"
	"encoding/json"
	"fmt"
	tpl "github.com/shieldworks/aegis/core/template"
	"strings"
	"text/template"
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

// parseForK8sSecret parses the provided `SecretStored` and applies a template
// if one is defined.
//
// Args:
//
//	secret: A SecretStored struct containing the secret data and metadata.
//
// Returns:
//
//	A map of string keys to string values, containing the parsed secret data.
//
//	If there is an error during parsing or applying the template, an error
//	will be returned.
func parseForK8sSecret(secret SecretStored) (map[string]string, error) {
	// cannot move this to /core/template because of circular dependency.

	jsonData := strings.TrimSpace(secret.Value)
	tmpStr := strings.TrimSpace(secret.Meta.Template)

	secretData := make(map[string]string)
	err := json.Unmarshal([]byte(jsonData), &secretData)
	if err != nil {
		return secretData, err
	}

	if tmpStr == "" {
		return secretData, err
	}

	tmpl, err := template.New("secret").Parse(tmpStr)
	if err != nil {
		return secretData, err
	}

	var tpl bytes.Buffer
	err = tmpl.Execute(&tpl, secretData)
	if err != nil {
		return secretData, err
	}

	output := make(map[string]string)
	err = json.Unmarshal(tpl.Bytes(), &output)
	if err != nil {
		return output, err
	}

	return output, nil
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
	newData, err := parseForK8sSecret(secret)
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

// Parse takes a data.SecretStored type as input and returns the parsed
// string or an error.
//
// If the Meta.Template field is empty, it tries to parse secret.Value;
// otherwise it transforms secret.Value using the Go template transformation
// defined by Meta.Template.
//
// If the Meta.Format field is None, it returns the parsed string.
//
// If the Meta.Format field is Json, it returns the parsed string if it’s a
// valid JSON or the original string otherwise.
//
// If the Meta.Format field is Yaml, it tries its best to transform the data
// into Yaml. If it fails, it tries to return a valid JSON at least. If that
// fails too, returns the original secret value.
//
// If the Meta.Format field is not recognized, it returns an empty string.
func (secret SecretStored) Parse() (string, error) {
	jsonData := strings.TrimSpace(secret.Value)
	tmpStr := strings.TrimSpace(secret.Meta.Template)

	parsedString := ""
	if tmpStr == "" {
		parsedString = jsonData
	} else {
		parsedString = tpl.TryParse(tmpStr, jsonData)
	}

	switch secret.Meta.Format {
	case None:
		return parsedString, nil
	case Json:
		if tpl.ValidJSON(parsedString) {
			return parsedString, nil
		} else {
			return jsonData, nil
		}
	case Yaml:
		if tpl.ValidJSON(parsedString) {
			yml, err := tpl.JsonToYaml(parsedString)
			if err != nil {
				return parsedString, err
			}
			return yml, nil
		} else {
			yml, err := tpl.JsonToYaml(jsonData)
			if err != nil {
				return jsonData, err
			}
			return yml, nil
		}
	}

	// Unknown option.
	return "", nil
}
