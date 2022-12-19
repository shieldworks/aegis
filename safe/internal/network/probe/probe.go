/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                  <aegis.z2h.dev>
 *     .\_/.
 */

package probe

import (
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
	res.WriteHeader(http.StatusOK)
	_, err := res.Write([]byte("OK"))
	if err != nil {
		log.Printf("problem sending response: %s", err.Error())
		return
	}
}
