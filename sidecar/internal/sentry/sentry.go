/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                  <aegis.z2h.dev>
 *     .\_/.
 */

package sentry

import (
	"aegis-sidecar/internal/state"
	"fmt"
	"time"
)

func fetchSecrets() {
	fmt.Println(state.Id(), state.Secret(), state.SafeApiRoot())
}

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
