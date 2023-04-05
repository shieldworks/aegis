/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                    <aegis.ist>
 *     .\_/.
 */

package probe

import (
	"github.com/shieldworks/aegis/core/env"
	"log"
	"net/http"
)

// CreateLiveness sets up and starts an HTTP server on the port specified by
// env.ProbeLivenessPort() to serve as a liveness probe for the application.
// The server listens for requests at the root path ("/") and responds with an
// "ok" message. If there is an error starting the server, the function logs
// a fatal message and returns.
func CreateLiveness() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", ok)
	err := http.ListenAndServe(env.ProbeLivenessPort(), mux)
	if err != nil {
		log.Fatalf("error creating liveness probe: %s", err.Error())
		return
	}
}

// CreateReadiness sets up and starts an HTTP server on the port specified by
// env.ProbeReadinessPort() to serve as a readiness probe for the application.
// The server listens for requests at the root path ("/") and responds with an
// "ok" message. If there is an error starting the server, the function logs
// a fatal message and returns.
func CreateReadiness() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", ok)
	err := http.ListenAndServe(env.ProbeReadinessPort(), mux)
	if err != nil {
		log.Fatalf("error creating readiness probe: %s", err.Error())
		return
	}
}
