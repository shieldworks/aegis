/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                  <aegis.z2h.dev>
 *     .\_/.
 */

package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println("hello from notary!")
		time.Sleep(5 * time.Second)
	}
}
