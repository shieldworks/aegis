/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                  <aegis.z2h.dev>
 *     .\_/.
 */

package v1

import (
	reqres "aegis-sidecar/internal/entity/reqres/v1"
	"context"
	"encoding/json"
	"log"
	"net/http"
)

func DecodeBootstrapRequest(
	_ context.Context, r *http.Request,
) (interface{}, error) {
	var request reqres.BootstrapRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		log.Printf("DecodeBootstrapRequest: error decoding: %s\n", err.Error())

		request.Err = "DecodeBootstrapRequest: Problem decoding JSON."
	}

	return request, nil
}
