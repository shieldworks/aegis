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

const secretFilePath = "/opt/aegis/secrets.json"

func main() {
	for {
		dat, err := os.ReadFile(secretFilePath)
		if err != nil {
			fmt.Println("Failed to read the secrets file. Will retry in 5 seconds…")
		} else {
			fmt.Println("secret: '", string(dat), "'")
		}

		time.Sleep(5 * time.Second)
	}
}
