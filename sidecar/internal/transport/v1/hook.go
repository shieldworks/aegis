/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                  <aegis.z2h.dev>
 *     .\_/.
 */

package v1

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
)

func DecodeHookRequest(
	_ context.Context, r *http.Request,
) (interface{}, error) {
	var request reqres.HookRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		log.Printf("DecodeHookRequest: error decoding: %s\n", err.Error())

		request.Err = "DecodeHookRequest: Problem decoding JSON."
	}

	return request, nil
}
