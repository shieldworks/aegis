/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                  <aegis.z2h.dev>
 *     .\_/.
 */

package handle

import (
	"aegis-safe/internal/server/route"
	"github.com/spiffe/go-spiffe/v2/svid/x509svid"
	"io"
	"log"
	"net/http"
)

func InitializeRoutes() {
	// Set up a `/` resource handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r == nil {
			return
		}

		tlsConnectionState := r.TLS
		if len(tlsConnectionState.PeerCertificates) == 0 {
			log.Println("no peer certs. exiting.")
			return
		}

		id, err := x509svid.IDFromCert(tlsConnectionState.PeerCertificates[0])
		if err != nil {
			log.Println("problem extracting svid. exiting.")
			return
		}

		sid := id.String()
		p := r.URL.Path

		// sidecar -> safe : fetch secrets
		if r.Method == http.MethodPost && p == "/v1/fetch" {
			route.Fetch(w, r, sid)
			return
		}

		// sentinel -> safe : put secrets
		if r.Method == http.MethodPost && p == "/v1/secret" {
			route.Secret(w, r, sid)
			return
		}

		w.WriteHeader(http.StatusBadRequest)
		_, err = io.WriteString(w, "")
		if err != nil {
			log.Println("Problem writing response")
			return
		}
	})
}
