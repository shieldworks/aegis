/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                    <aegis.ist>
 *     .\_/.
 */

package handle

import (
	"github.com/pkg/errors"
	"github.com/shieldworks/aegis/app/safe/internal/server/route"
	"github.com/shieldworks/aegis/core/log"
	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/go-spiffe/v2/svid/x509svid"
	"io"
	"net/http"
)

func getSpiffeId(r *http.Request) (*spiffeid.ID, error) {
	tlsConnectionState := r.TLS
	if len(tlsConnectionState.PeerCertificates) == 0 {
		return nil, errors.New("no peer certs")
	}

	id, err := x509svid.IDFromCert(tlsConnectionState.PeerCertificates[0])
	if err != nil {
		return nil, errors.Wrap(err, "problem extracting svid")
	}

	return &id, nil
}

func InitializeRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		id, err := getSpiffeId(r)
		if err != nil {
			log.WarnLn("Handler: blocking insecure svid", id, err)

			// Block insecure connection attempt.
			_, err = io.WriteString(w, "")
			if err != nil {
				log.InfoLn("Problem writing response:", err.Error())
				return
			}
		}

		sid := id.String()
		p := r.URL.Path

		log.DebugLn("Handler: got svid:", sid, "path", p, "method", r.Method)

		// Route to list secrets.
		// Only Aegis Sentinel is allowed to call this API endpoint.
		// Calling it from anywhere else will error out.
		if r.Method == http.MethodGet && p == "/sentinel/v1/secrets" {
			log.DebugLn("Handler: will list")
			route.List(w, r, sid)
			return
		}

		// Route to add secrets to Aegis Safe.
		// Only Aegis Sentinel is allowed to call this API endpoint.
		// Calling it from anywhere else will error out.
		if r.Method == http.MethodPost && p == "/sentinel/v1/secrets" {
			log.DebugLn("Handler: will secret")
			route.Secret(w, r, sid)
			return
		}

		// Route to fetch secrets.
		// Only an Aegis-nominated workload is allowed to
		// call this API endpoint. Calling it from anywhere else will
		// error out.
		if r.Method == http.MethodGet && p == "/workload/v1/secrets" {
			log.DebugLn("Handler: will fetch")
			route.Fetch(w, r, sid)
			return
		}

		log.DebugLn("Handler: route mismatch")

		w.WriteHeader(http.StatusBadRequest)
		_, err = io.WriteString(w, "")
		if err != nil {
			log.WarnLn("Problem writing response", err.Error())
			return
		}
	})
}
