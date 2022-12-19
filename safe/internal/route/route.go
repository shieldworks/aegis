/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                  <aegis.z2h.dev>
 *     .\_/.
 */

package route

import (
	"aegis-safe/internal/network/probe"
	v1Network "aegis-safe/internal/network/v1"
	"github.com/gorilla/mux"
	"net/http"
)

func Probes(r *mux.Router) {
	// Liveness endpoint.
	r.Methods(http.MethodGet).Path("/healthz").HandlerFunc(probe.Health)

	// Readiness endpoint.
	// Will fail if Safe is not bootstrapped.
	r.Methods(http.MethodGet).Path("/readyz").HandlerFunc(probe.Ready)
}

func AdminEndpoints(r *mux.Router, api *v1Network.Api) {
	// Shall return an error if Safe is not bootstrapped.
	// Only administrator can use this method.
	r.Methods(http.MethodPut).Path("/v1/secret").Handler(api.SecretUpsert)

	// TODO: we probably don’t need this. If so, remove all trails.
	//// to register workloads to notary.
	//// requires bootstrap token.
	//r.Methods(http.MethodPut).Path("/v1/register")
}

func WorkloadEndpoints(r *mux.Router, api *v1Network.Api) {
	// TODO: only sidecar can read this with a proper token.
	// shall return an error if safe is not bootstrapped.
	r.Methods(http.MethodGet).Path("/v1/fetch").Handler(api.SecretFetch)
}

func NotaryEndpoints(r *mux.Router, api *v1Network.Api) {
	// This will be triggered from `notary`. The `AEGIS_NOTARY_ID` environment
	// variable that is passed in the payload, should match the id that `safe`
	// is initialized with for the method to succeed.
	r.Methods(http.MethodPost).Path("/v1/bootstrap").Handler(api.Bootstrap)

	// hook to register workload keys
	// Only notary can call this; to call, it needs the bootstrap token.
	r.Methods(http.MethodPut).Path("/v1/workload").Handler(api.WorkloadRegister)
}
