/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                  <aegis.z2h.dev>
 *     .\_/.
 */

package main

import (
	"aegis-safe/internal/env"
	v1Network "aegis-safe/internal/network/v1"
	"aegis-safe/internal/route"
	v1Service "aegis-safe/internal/service/v1"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	apiV1 := &v1Network.Api{}

	r := mux.NewRouter()

	// Bind handlers.
	v1Network.Init(apiV1, v1Service.NewApiV1Service())

	// Bind other routes.
	route.Probes(r)
	route.AdminEndpoints(r, apiV1)
	route.SidecarEndpoints(r, apiV1)
	route.NotaryEndpoints(r, apiV1)

	p, a := env.Port(), env.AppName()
	log.Printf("[SAFE]: '%s' will listen at port '%s'.", a, p)
	log.Fatal(http.ListenAndServe(p, r))
}
