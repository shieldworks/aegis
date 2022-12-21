/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                  <aegis.z2h.dev>
 *     .\_/.
 */

package route

import (
	v1Network "aegis-sidecar/internal/network/v1"
	"github.com/gorilla/mux"
	"net/http"
)

func HookEndpoints(r *mux.Router, api *v1Network.Api) {
	r.Methods(http.MethodPut).Path("/v1/hook").Handler(api.Hook)
}
