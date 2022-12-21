/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                  <aegis.z2h.dev>
 *     .\_/.
 */

package sentry

import (
	"time"
)

// Watch synchronizes the internal state of the sidecar by talking to
// `safe` regularly.
func Watch() {
	// TODO: make this configurable.
	ticker := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-ticker.C:
			fetchSecrets()
		}
	}
}
