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
	endpoint "aegis-safe/internal/endpoint/v1/secret"
	service "aegis-safe/internal/service/v1"
	transport "aegis-safe/internal/transport/v1"
	"net/http"
)

type SecretHandlers struct {
	SecretUpsert http.Handler
	SecretFetch  http.Handler
}

func DefineSecretHandlers(s *Api, svc service.ApiV1Service) {
	s.SecretUpsert = coreHttp.Serve(
		endpoint.MakeSecretUpsertEndpoint(svc),
		transport.DecodeSecretUpsertRequest,
		transport.EncodeResponse,
	)

	s.SecretFetch = coreHttp.Serve(
		endpoint.MakeSecretFetchEndpoint(svc),
		transport.DecodeSecretFetchRequest,
		transport.EncodeResponse,
	)
}
