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
	"os"
	"os/signal"
	"syscall"
	"time"
)

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
		fmt.Printf("My secret: '%s'.\n", os.Getenv("SECRET"))
		fmt.Printf("My creds: username:'%s' password:'%s'.\n",
			os.Getenv("USERNAME"), os.Getenv("PASSWORD"),
		)
		fmt.Println("")

		time.Sleep(5 * time.Second)
	}
}
