/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                    <aegis.ist>
 *     .\_/.
 */

package main

import (
	"fmt"
	"github.com/shieldworks/aegis/sdk/sentry"
)

func main() {
	d, err := sentry.Fetch()
	if err != nil {
		fmt.Print("ERR_SENTRY_FETCH_FAILED")
		return
	}

	if d.Data == "" {
		fmt.Print("NO_SECRET")
	}

	fmt.Print(d.Data)
}
