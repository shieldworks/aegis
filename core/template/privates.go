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
	"gopkg.in/yaml.v3"
	"text/template"
)

func validJSON(s string) bool {
	var js map[string]interface{}
	return json.Unmarshal([]byte(s), &js) == nil
}

func jsonToYaml(js string) (string, error) {
	var jsonObj interface{}
	err := json.Unmarshal([]byte(js), &jsonObj)
	if err != nil {
		return "", err
	}
	yamlBytes, err := yaml.Marshal(jsonObj)
	if err != nil {
		return "", err
	}
	return string(yamlBytes), nil
}

func tryParse(tmpStr, json string) string {
	tmpl, err := template.New("secret").Parse(tmpStr)
	if err != nil {
		return json
	}

	var tpl bytes.Buffer
	err = tmpl.Execute(&tpl, json)
	if err != nil {
		return json
	}

	return tpl.String()
}
