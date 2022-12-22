/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                  <aegis.z2h.dev>
 *     .\_/.
 */

package sentry

import (
	reqres "aegis-sidecar/internal/entity/reqres/v1"
	"aegis-sidecar/internal/state"
	"bufio"
	"context"
	"fmt"
	"os"
)

func saveData(data string) {
	path := "/opt/aegis/secrets.json"

	// fmt.Println("path:", path)

	f, err := os.Create(path)
	if err != nil {
		// TODO: handle me.
		panic("poop!")
	}

	w := bufio.NewWriter(f)
	n, err := w.WriteString(data)
	if err != nil {
		// TODO: handle me
		panic("poop!")
	}

	// fmt.Println("wrote", n, "bytes.")

	err = w.Flush()
	if err != nil {
		// TODO: handle
		panic("poop")
	}
}

func fetchSecrets() {
	if !state.Bootstrapped() {
		return
	}

	id := state.Id()
	secret := state.Secret()

	fmt.Println(state.Id(), state.Secret(), state.SafeApiRoot())

	res, err := newSafeFetchEndpoint()(
		context.Background(),
		reqres.SecretFetchRequest{
			WorkloadId:     id,
			WorkloadSecret: secret,
		})
	if err != nil {
		// TODO: handle me
		panic("handle me")
	}

	sfr, ok := res.(reqres.SecretFetchResponse)
	if !ok {
		// TODO: handle me
		panic("handle me!")
	}

	data := sfr.Data

	// TODO: save data to /opt/aegis/secrets.json
	// TODO: make the filename configurable.
	// fmt.Println("data: '", data, "'")

	saveData(data)
}
