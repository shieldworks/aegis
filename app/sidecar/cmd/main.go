/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                    <aegis.ist>
 *     .\_/.
 */

package main

import (
	"github.com/shieldworks/aegis/core/log"
	"github.com/shieldworks/aegis/core/system"
	"github.com/shieldworks/aegis/sdk/sentry"
)

func main() {
	id := "AEGSSDCR"
	log.InfoLn(&id, "Starting Aegis Sidecar")
	go sentry.Watch()
	// Keep the main routine alive:
	system.KeepAlive()
}
