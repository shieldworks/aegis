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
	"os"
	"os/signal"
	"syscall"
)

type Secret struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func main() {
	go func() {
		// Block the process from exiting, but also be graceful and honor the
		// termination signals that may come from the orchestrator.
		s := make(chan os.Signal, 1)
		signal.Notify(s, syscall.SIGINT, syscall.SIGTERM)
		select {
		case e := <-s:
			fmt.Println(e)
			panic("bye cruel world!")
		}
	}()

	// Fetch the secret from the Aegis Safe.
	d, err := sentry.Fetch()
	if err != nil {
		fmt.Println("Failed to fetch the secrets. Try again later.")
		fmt.Println(err.Error())
		return
	}

	if d.Data == "" {
		fmt.Println("No secret yet… Try again later.")
		return
	}

	// d.Data is a collection of Secrets.
	fmt.Println(d.Data)
}
