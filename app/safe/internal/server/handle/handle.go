/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                    <aegis.ist>
 *     .\_/.
 */

package handle

import (
	"github.com/shieldworks/aegis/app/safe/internal/server/route"
	"github.com/shieldworks/aegis/core/crypto"
	"github.com/shieldworks/aegis/core/log"
	"io"
	"net/http"
)

func InitializeRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		cid, _ := crypto.RandomString(8)
		if cid == "" {
			cid = "AEGISFHN"
		}

		id, err := spiffeIdFromRequest(r)
		if err != nil {
			log.WarnLn(&cid, "Handler: blocking insecure svid", id, err)

			// Block insecure connection attempt.
			_, err = io.WriteString(w, "")
			if err != nil {
				log.InfoLn(&cid, "Problem writing response:", err.Error())
				return
			}
		}

		sid := id.String()
		p := r.URL.Path

		log.DebugLn(&cid, "Handler: got svid:", sid, "path", p, "method", r.Method)

		// Route to list secrets.
		// Only Aegis Sentinel is allowed to call this API endpoint.
		// Calling it from anywhere else will error out.
		if r.Method == http.MethodGet && p == "/sentinel/v1/secrets" {
			log.DebugLn(&cid, "Handler: will list")
			route.List(cid, w, r, sid)
			return
		}

		// Route to define the master key when AEGIS_MANUAL_KEY_INPUT is set.
		// Only Aegis Sentinel is allowed to call this API endpoint.
		// This method works only once. Once a key is set, there is no way to
		// update it. You will have to kill the Aegis Sentinel pod and restart it
		// to be able to set a new key.
		if r.Method == http.MethodPost && p == "/sentinel/v1/keys" {
			log.DebugLn(&cid, "Handler: will receive keys")
			route.ReceiveKeys(cid, w, r, sid)
			return
		}

		// Route to add secrets to Aegis Safe.
		// Only Aegis Sentinel is allowed to call this API endpoint.
		// Calling it from anywhere else will error out.
		if r.Method == http.MethodPost && p == "/sentinel/v1/secrets" {
			log.DebugLn(&cid, "Handler:/sentinel/v1/secrets will secret")
			route.Secret(cid, w, r, sid)
			return
		}

		// Route to delete secrets from Aegis Safe.
		// Only Aegis Sentinel is allowed to call this API endpoint.
		// Calling it from anywhere else will error out.
		if r.Method == http.MethodDelete && p == "/sentinel/v1/secrets" {
			log.DebugLn(&cid, "Handler:/sentinel/v1/secrets will delete")
			route.Delete(cid, w, r, sid)
			return
		}

		// Route to fetch secrets.
		// Only an Aegis-nominated workload is allowed to
		// call this API endpoint. Calling it from anywhere else will
		// error out.
		if r.Method == http.MethodGet && p == "/workload/v1/secrets" {
			log.DebugLn(&cid, "Handler:/workload/v1/secrets: will fetch")
			route.Fetch(cid, w, r, sid)
			return
		}

		log.DebugLn(&cid, "Handler: route mismatch")

		w.WriteHeader(http.StatusBadRequest)
		_, err = io.WriteString(w, "")
		if err != nil {
			log.WarnLn(&cid, "Problem writing response", err.Error())
			return
		}
	})
}
