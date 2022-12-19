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
	endpoint "aegis-safe/internal/endpoint/v1/bootstrap"
	service "aegis-safe/internal/service/v1"
	transport "aegis-safe/internal/transport/v1"
	"net/http"
)

type BootstrapHandlers struct {
	Bootstrap http.Handler
}

func DefineBootstrapHandlers(s *Api, svc service.ApiV1Service) {
	s.Bootstrap = coreHttp.Serve(
		endpoint.MakeBootstrapEndpoint(svc),
		transport.DecodeBootstrapRequest,
		transport.EncodeResponse,
	)
}
