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

func main() {
	id := os.Getenv("AEGIS_ID")
	secret := os.Getenv("AEGIS_SECRET")

	for {
		fmt.Printf("[AEGIS-SIDECAR]: will fetch (%s) of (%s).\n", secret, id)
		time.Sleep(5 * time.Second)
	}
}
