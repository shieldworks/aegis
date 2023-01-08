/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                  <aegis.z2h.dev>
 *     .\_/.
 */

package route

import (
	v1 "aegis-safe/internal/entity/reqres/v1"
	"aegis-safe/internal/state"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strings"
)

func Fetch(w http.ResponseWriter, r *http.Request, svid string) {
	if r == nil {
		return
	}

	if !strings.HasPrefix(svid, "spiffe://aegis.z2h.dev/workload/") {
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

	var sr v1.SecretFetchRequest

	err = json.Unmarshal(body, &sr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err := io.WriteString(w, "")
		if err != nil {
			log.Println("Problem sending response")
		}
		return
	}

	tmp := strings.Replace(svid, "spiffe://aegis.z2h.dev/workload/", "", 1)
	parts := strings.Split(tmp, "/")
	if len(parts) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		_, err := io.WriteString(w, "")
		if err != nil {
			log.Println("Problem sending response")
		}
		return
	}

	workloadId := parts[0]
	value := state.ReadSecret(workloadId)

	sfr := v1.SecretFetchResponse{
		Data: value,
	}

	resp, err := json.Marshal(sfr)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := io.WriteString(w, "Problem unmarshaling response")
		if err != nil {
			log.Println("Problem sending response")
		}
		return
	}

	_, err = io.WriteString(w, string(resp))
	if err != nil {
		log.Println("Problem sending response")
	}
}
