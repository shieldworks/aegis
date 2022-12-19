/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                  <aegis.z2h.dev>
 *     .\_/.
 */

package probe

import (
	"aegis-safe/internal/state"
	"log"
	"net/http"
)

func Health(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)
	_, err := res.Write([]byte("OK"))
	if err != nil {
		log.Printf("problem sending response: %s", err.Error())
		return
	}
}

func Ready(res http.ResponseWriter, req *http.Request) {
	bootstrapped := state.Bootstrapped()

	if !bootstrapped {
		res.WriteHeader(http.StatusServiceUnavailable)
		_, err := res.Write([]byte("Safe has not bootstrapped yet."))
		if err != nil {
			log.Printf("problem sending response: %s", err.Error())
			return
		}
		return
	}

	res.WriteHeader(http.StatusOK)
	_, err := res.Write([]byte("OK"))
	if err != nil {
		log.Printf("problem sending response: %s", err.Error())
		return
	}
}
