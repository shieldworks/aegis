/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                  <aegis.z2h.dev>
 *     .\_/.
 */

package route

import (
	"aegis-safe/internal/state"
	"encoding/json"
	"github.com/zerotohero-dev/aegis-core/entity/reqres/v1"
	"github.com/zerotohero-dev/aegis-core/validation"
	"io"
	"log"
	"net/http"
)

func Secret(w http.ResponseWriter, r *http.Request, svid string) {
	if r == nil {
		return
	}

	if !validation.IsSentinel(svid) {
		w.WriteHeader(http.StatusBadRequest)
		_, err := io.WriteString(w, "")
		if err != nil {
			log.Println("Problem sending response")
		}
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err := io.WriteString(w, "")
		if err != nil {
			log.Println("Problem sending response")
		}
		return
	}
	defer func(b io.ReadCloser) {
		if b == nil {
			return
		}
		err := b.Close()
		if err != nil {
			log.Println("Problem closing body")
		}
	}(r.Body)

	var sr v1.SecretUpsertRequest

	err = json.Unmarshal(body, &sr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err := io.WriteString(w, "")
		if err != nil {
			log.Println("Problem sending response")
		}
		return
	}

	workloadId := sr.WorkloadId
	value := sr.Value

	state.UpsertSecret(workloadId, value)

	_, err = io.WriteString(w, "OK")
	if err != nil {
		log.Println("Problem sending response")
	}
}
