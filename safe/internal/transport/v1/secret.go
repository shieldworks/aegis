/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                  <aegis.z2h.dev>
 *     .\_/.
 */

package v1

import (
	"aegis-safe/internal/entity/reqres"
	"context"
	"encoding/json"
	"log"
	"net/http"
)

func DecodeSecretUpsertRequest(
	_ context.Context, r *http.Request,
) (interface{}, error) {
	var request reqres.SecretUpsertRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		log.Printf("DecodeUpsertSecretRequest: error decoding: %s\n", err.Error())

		request.Err = "DecodeUpsertSecretRequest: Problem decoding JSON."
	}

	return request, nil
}

func DecodeSecretFetchRequest(
	_ context.Context, r *http.Request,
) (interface{}, error) {
	var request reqres.SecretFetchRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		log.Printf("DecodeSecretFetchRequest: error decoding: %s\n", err.Error())

		request.Err = "DecodeSecretFetchRequest: Problem decoding JSON."
	}

	return request, nil
}
