/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                    <aegis.ist>
 *     .\_/.
 */

package v1

import (
	"fmt"
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
