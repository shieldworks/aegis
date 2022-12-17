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
	"os"
	"time"
)

func main() {
	fmt.Println("Hello world, hello stars, hello universe!")

	const secretFilePath = "/opt/aegis/secrets.json"

	for {
		dat, err := os.ReadFile(secretFilePath)
		if err != nil {
			panic("Failed to read secrets file.")
		}
		fmt.Print(string(dat))

		time.Sleep(5 * time.Second)
	}
}
