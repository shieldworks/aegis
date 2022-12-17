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
	"time"
)

func main() {

	// /healthz
	// /readyz
	// POST /v1/secret/podName:podLabel (only for admin)
	// GET /v1/secret/podName:podLabel (mTLS: only for sidecar)

	for {
		fmt.Println("hello from safe!")
		time.Sleep(5 * time.Second)
	}
}
