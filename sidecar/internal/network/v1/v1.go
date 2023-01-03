/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                  <aegis.z2h.dev>
 *     .\_/.
 */

package v1

import service "aegis-sidecar/internal/service/v1"

type Api struct {
	HookHandlers
}

func Init(s *Api, svc service.ApiV1Service) {
	// DefineHookHandlers(s, svc)
}
