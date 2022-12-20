/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                  <aegis.z2h.dev>
 *     .\_/.
 */

package v1

import (
	coreHttp "aegis-safe/internal/core/http"
	reqres "aegis-safe/internal/entity/reqres/v1"
	"context"
	"encoding/json"
	"log"
	"net/http"
)

func EncodeResponse(
	_ context.Context, w http.ResponseWriter, response interface{},
) error {
	responseErr := coreHttp.Err(response)
	if responseErr != "" {
		log.Printf("EncodeResponse: error encoding response: %s\n", responseErr)

		res := reqres.GenericResponse{
			Err: "There is a problem in your request.",
		}

		w.WriteHeader(http.StatusBadRequest)

		return json.NewEncoder(w).Encode(res)
	}

	return json.NewEncoder(w).Encode(response)
}
