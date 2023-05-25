/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                    <aegis.ist>
 *     .\_/.
 */

package main

// Note that there is no Aegis-specific dependency in the app’s code:
// That’s the benefit of using “Aegis Sidecar”: The application has zero idea
// that `Aegis Safe` exists. From its perspective, it just knows that there
// are secrets in a predefined location that it can read and parse. And, that’s
// a good way to separate cross-cutting concerns.

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func sidecarSecretsPath() string {
	p := os.Getenv("AEGIS_SIDECAR_SECRETS_PATH")
	if p == "" {
		p = "/opt/aegis/secrets.json"
	}
	return p
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

	for {
		dat, err := os.ReadFile(sidecarSecretsPath())
		if err != nil {
			fmt.Println("Failed to read the secrets file. Will retry in 5 seconds…")
			fmt.Println(err.Error())
		} else {
			fmt.Println("secret: '", string(dat), "'")
		}

		time.Sleep(5 * time.Second)
	}
}
