/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                  <aegis.z2h.dev>
 *     .\_/.
 */

package route

import "net/http"

func HookEndpoints(r *mux.Router, api *v1Network.Api) {
	r.Methods(http.MethodPut).Path("/v1/hook").Handler(api.HookSync)
}
