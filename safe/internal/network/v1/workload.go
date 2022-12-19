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
	endpoint "aegis-safe/internal/endpoint/v1/workload"
	service "aegis-safe/internal/service/v1"
	transport "aegis-safe/internal/transport/v1"
	"net/http"
)

type WorkloadHandlers struct {
	WorkloadRegister http.Handler
}

func DefineWorkloadHandlers(s *Api, svc service.ApiV1Service) {
	s.WorkloadRegister = coreHttp.Serve(
		endpoint.MakeWorkloadRegisterEndpoint(svc),
		transport.DecodeWorkloadRegisterRequest,
		transport.EncodeResponse,
	)
}
