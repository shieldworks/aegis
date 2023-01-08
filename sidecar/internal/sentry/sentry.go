/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                  <aegis.z2h.dev>
 *     .\_/.
 */

package sentry

import (
	"aegis-sidecar/internal/env"
	"time"
)

// Watch synchronizes the internal state of the sidecar by talking to
// `safe` regularly.
func Watch() {
	ticker := time.NewTicker(env.SentryPollInterval())
	for {
		select {
		case <-ticker.C:
			fetchSecrets()
		}
	}
}
