/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                    <aegis.ist>
 *     .\_/.
 */

package main

import (
	"github.com/shieldworks/aegis/core/system"
)

func main() {
	// Run on the main thread to wait forever.
	system.KeepAlive()
}
