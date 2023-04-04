/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                    <aegis.ist>
 *     .\_/.
 */

package template

import (
	"bytes"
	"encoding/json"
	data "github.com/shieldworks/aegis/core/entity/data/v1"
	"strings"
	"text/template"
)

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
func Parse(secret data.SecretStored) (string, error) {
	jsonData := strings.TrimSpace(secret.Value)
	tmpStr := strings.TrimSpace(secret.Meta.Template)

	parsedString := ""
	if tmpStr == "" {
		parsedString = jsonData
	} else {
		parsedString = tryParse(tmpStr, jsonData)
	}

	switch secret.Meta.Format {
	case data.None:
		return parsedString, nil
	case data.Json:
		if validJSON(parsedString) {
			return parsedString, nil
		} else {
			return jsonData, nil
		}
	case data.Yaml:
		if validJSON(parsedString) {
			yml, err := jsonToYaml(parsedString)
			if err != nil {
				return parsedString, err
			}
			return yml, nil
		} else {
			yml, err := jsonToYaml(jsonData)
			if err != nil {
				return jsonData, err
			}
			return yml, nil
		}
	}

	// Unknown option.
	return "", nil
}

// ParseForK8sSecret parses the provided `SecretStored` and applies a template
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
func ParseForK8sSecret(secret data.SecretStored) (map[string]string, error) {
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
