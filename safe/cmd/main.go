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

	// Liveness endpoint.
	r.Methods(http.MethodGet).Path("/healthz").HandlerFunc(probe.Health)

	// Readiness endpoint.
	// Will fail if Safe is not bootstrapped.
	r.Methods(http.MethodGet).Path("/readyz").HandlerFunc(probe.Ready)

	// Shall return an error if Safe is not bootstrapped.
	// Only administrator can use this method.
	r.Methods(http.MethodPut).Path("/v1/secret/{value}").Handler(apiV1.SecretUpsert)

	// TODO: only sidecar can read this with a proper token.
	// shall return an error if safe is not bootstrapped.
	r.Methods(http.MethodGet).Path("/v1/fetch").Handler(apiV1.SecretFetch)

	// TODO: implement me.
	// TODO: shall be triggered from notary. Safe is not ready until it is
	// bootstrapped.
	r.Methods(http.MethodPost).Path("/bootstrap")

	// hook to register workload keys
	// Only notary can call this; to call it needs the bootstrap key.
	r.Methods(http.MethodPut).Path("/v1/workload")

	// to register workloads to notary.
	// requires bootstrap key.
	r.Methods(http.MethodPut).Path("/v1/register")

	port := ":8017"

	log.Printf("[SAFE]: '%s' will listen at port '%s'.", appName, port)
	log.Fatal(http.ListenAndServe(port, r))
}
