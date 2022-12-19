/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                  <aegis.z2h.dev>
 *     .\_/.
 */

package v1

import service "aegis-safe/internal/service/v1"

type Api struct {
	SecretHandlers
	BootstrapHandlers
	WorkloadHandlers
}

func Init(s *Api, svc service.ApiV1Service) {
	DefineSecretHandlers(s, svc)
	DefineBootstrapHandlers(s, svc)
	DefineWorkloadHandlers(s, svc)
}
