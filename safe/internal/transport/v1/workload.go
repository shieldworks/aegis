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

func DecodeWorkloadRegisterRequest(
	_ context.Context, r *http.Request,
) (interface{}, error) {
	var request reqres.WorkloadRegisterRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		log.Printf("DecodeWorkloadRegisterRequest: error decoding: %s\n", err.Error())

		request.Err = "DecodeWorkloadRegisterRequest: Problem decoding JSON."
	}

	return request, nil
}
