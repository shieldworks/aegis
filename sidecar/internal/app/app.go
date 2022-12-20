/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                  <aegis.z2h.dev>
 *     .\_/.
 */

package app

import (
	"aegis-sidecar/internal/route"
	"log"
	"net/http"
)

func Serve() {
	// web server at localhost: 8039
	// /aegis/v1/hook => { id, token, safeApiRoot }

	apiV1 := &v1Network.Api{}

	r := mux.NewRouter()

	// Bind handlers.
	v1Network.Init(apiV1, v1Service.NewApiV1Service())

	route.HookEndpoints(r, apiV1)

	p, a := env.Port(), env.AppName()

	log.Printf("[SIDECAR]: '%s' will listen at port '%s'.", a, p)
	log.Fatal(http.ListenAndServe(p, r))
}
