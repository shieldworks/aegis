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
	"strings"
)

func main() {
	d, err := sentry.Fetch()
	if err != nil {
		msg := err.Error()

		if strings.Contains(msg, "Secret does not exist") {
			fmt.Print("NO_SECRET")
			return
		}

		fmt.Print("ERR_SENTRY_FETCH_FAILED")
		fmt.Print(" ", err.Error())
		return
	}

	if strings.TrimSpace(d.Data) == "" {
		fmt.Print("NO_SECRET")
	}

	fmt.Print(d.Data)
}
