/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                    <aegis.ist>
 *     .\_/.
 */

package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/shieldworks/aegis/sdk/sentry"
	"os"
	"os/signal"
	"syscall"
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

	// Check if d.Data is a JSON array
	if string(d.Data[0]) == "[" {
		// Convert the array into a slice of strings
		var dataSlice []string
		err = json.Unmarshal([]byte(d.Data), &dataSlice)
		if err != nil {
			fmt.Println("Failed to unmarshal the data into a slice of strings. Check the data format.")
			fmt.Println(err.Error())
			return
		}

		// Concatenate all members of the slice into one large string
		concatString := ""
		for _, s := range dataSlice {
			concatString += s
		}

		// Base64 decode the string
		decodedString, err := base64.StdEncoding.DecodeString(concatString)
		if err != nil {
			fmt.Println("Failed to decode the base64 string.")
			fmt.Println(err.Error())
			return
		}

		// Print the result
		fmt.Println(string(decodedString))
	} else {
		// d.Data is a collection of Secrets.
		fmt.Println(d.Data)
	}
}
