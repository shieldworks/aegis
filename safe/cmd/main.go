/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                  <aegis.z2h.dev>
 *     .\_/.
 */

package main

import (
	"aegis-safe/internal/network/probe"
	v1Network "aegis-safe/internal/network/v1"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const appName = "aegis-safe"

func main() {
	apiV1 := &v1Network.Api{}

	r := mux.NewRouter()

	r.Methods(http.MethodGet).Path("/healthz").HandlerFunc(probe.Health)
	r.Methods(http.MethodGet).Path("/readyz").HandlerFunc(probe.Ready)

	r.Methods(http.MethodPut).Path("/v1/secret/{value}").Handler(apiV1.SecretUpsert)

	// TODO: this shall only be reachable via the sidecar thru mTLS.
	r.Methods(http.MethodGet).Path("/v1/secret/{value}").Handler(apiV1.SecretRead)

	port := ":8017"

	log.Printf("[SAFE]: '%s' will listen at port '%s'.", appName, port)
	log.Fatal(http.ListenAndServe(port, r))
}
