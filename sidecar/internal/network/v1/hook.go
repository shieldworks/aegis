/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                  <aegis.z2h.dev>
 *     .\_/.
 */

package v1

import (
	coreHttp "aegis-sidecar/internal/core/http"
	endpoint "aegis-sidecar/internal/endpoint/v1/hook"
	service "aegis-sidecar/internal/service/v1"
	transport "aegis-sidecar/internal/transport/v1"
	"net/http"
)

type HookHandlers struct {
	Hook http.Handler
}

func DefineHookHandlers(s *Api, svc service.ApiV1Service) {
	s.Hook = coreHttp.Serve(
		endpoint.MakeHookEndpoint(svc),
		transport.DecodeHookRequest,
		transport.EncodeResponse,
	)
}
