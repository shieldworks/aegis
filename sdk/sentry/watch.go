/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                    <aegis.ist>
 *     .\_/.
 */

package sentry

import (
	"github.com/shieldworks/aegis/core/log"
	"github.com/shieldworks/aegis/sdk/internal/timer"
	"time"
)

// Watch synchronizes the internal state of the sidecar by talking to
// Aegis Safe regularly. It periodically calls Fetch behind-the-scenes to
// get its work done. Once it fetches the secrets, it saves it to
// the location defined in the `AEGIS_SIDECAR_SECRETS_PATH` environment
// variable (`/opt/aegis/secrets.json` by default).
func Watch() {
	interval := timer.InitialInterval
	successCount := int64(0)
	errorCount := int64(0)

	for {
		ticker := time.NewTicker(interval)
		select {
		case <-ticker.C:
			err := fetchSecrets()

			// Update parameters based on success/failure.
			interval, successCount, errorCount = timer.ExponentialBackoff(
				err == nil, interval, successCount, errorCount,
			)

			if err != nil {
				log.InfoLn("Could not fetch secrets", err.Error(),
					". Will retry in", interval, ".")
			}

			ticker.Stop()
		}
	}
}
