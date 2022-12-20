/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                  <aegis.z2h.dev>
 *     .\_/.
 */

package main

import (
	"aegis-sidecar/internal/app"
	"aegis-sidecar/internal/sentry"
)

func main() {
	go app.Serve()
	go sentry.Watch()
	select {} // block forever
}
