/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                  <aegis.z2h.dev>
 *     .\_/.
 */

package main

import (
	"fmt"
	"os"
	"time"
)

func sidecarSecretsPath() string {
	p := os.Getenv("AEGIS_SIDECAR_SECRETS_PATH")
	if p == "" {
		p = "/opt/aegis/secrets.json"
	}
	return p
}

func main() {
	for {
		dat, err := os.ReadFile(sidecarSecretsPath())
		if err != nil {
			fmt.Println("Failed to read the secrets file. Will retry in 5 seconds…")
		} else {
			fmt.Println("secret: '", string(dat), "'")
		}

		time.Sleep(5 * time.Second)
	}
}
