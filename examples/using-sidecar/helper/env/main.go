/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                    <aegis.ist>
 *     .\_/.
 */

package main

import (
	"fmt"
	"os"
)

func sidecarSecretsPath() string {
	p := os.Getenv("AEGIS_SIDECAR_SECRETS_PATH")
	if p == "" {
		p = "/opt/aegis/secrets.json"
	}
	return p
}

func main() {
	dat, err := os.ReadFile(sidecarSecretsPath())
	if err != nil {
		fmt.Print("ERR_READ_SECRET")
	} else {
		fmt.Print(string(dat))
	}
}
