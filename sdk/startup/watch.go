/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                    <aegis.ist>
 *     .\_/.
 */

package startup

import (
	"github.com/shieldworks/aegis/core/crypto"
	"github.com/shieldworks/aegis/core/env"
	"github.com/shieldworks/aegis/core/log"
	"github.com/shieldworks/aegis/sdk/sentry"
	"os"
	"time"
)

func initialized() bool {
	r, _ := sentry.Fetch()
	v := r.Data
	return v != ""
}

// Watch continuously polls the associated secret of the workload to exist.
// If the secret exists, and it is not empty, the function exits the init
// container with a success status code (0).
func Watch() {
	interval := env.InitContainerPollInterval()
	ticker := time.NewTicker(interval)

	cid, _ := crypto.RandomString(8)
	if cid == "" {
		cid = "AEGISSDK"
	}

	for {
		select {
		case <-ticker.C:
			log.InfoLn(&cid, "init:: tick")
			if initialized() {
				log.InfoLn(&cid, "initialized… exiting the init process")
				os.Exit(0)
			}
		}
	}
}
