/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                    <aegis.ist>
 *     .\_/.
 */

package probe

import (
	"fmt"
	"github.com/shieldworks/aegis/core/env"
	"log"
	"net/http"
)

func ok(w http.ResponseWriter, _ *http.Request) {
	_, err := fmt.Fprintf(w, "OK")
	if err != nil {
		log.Printf("probe response failure: %s", err.Error())
		return
	}
}

func CreateLiveness() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", ok)
	err := http.ListenAndServe(env.ProbeLivenessPort(), mux)
	if err != nil {
		log.Fatalf("error creating liveness probe: %s", err.Error())
		return
	}
}

func CreateReadiness() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", ok)
	err := http.ListenAndServe(env.ProbeReadinessPort(), mux)
	if err != nil {
		log.Fatalf("error creating readiness probe: %s", err.Error())
		return
	}
}
