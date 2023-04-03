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
	"github.com/shieldworks/aegis/sdk/startup"
)

func main() {
	id := "AEGINCTR"

	log.InfoLn(&id, "Starting Aegis Init Container")
	go startup.Watch()

	// Block the process from exiting, but also be graceful and honor the
	// termination signals that may come from the orchestrator.
	system.KeepAlive()
}
