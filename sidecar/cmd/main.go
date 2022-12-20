/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                  <aegis.z2h.dev>
 *     .\_/.
 */

package main

import (
	"aegis-demo-workload/internal/app"
	"aegis-demo-workload/internal/sentry"
)

func main() {
	go app.Serve()
	go sentry.Watch()
	select {} // block forever
}
