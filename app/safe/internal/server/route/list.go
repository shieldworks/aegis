/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                    <aegis.ist>
 *     .\_/.
 */

package route

import (
	"encoding/json"
	"github.com/shieldworks/aegis/app/safe/internal/state"
	"github.com/shieldworks/aegis/core/audit"
	reqres "github.com/shieldworks/aegis/core/entity/reqres/safe/v1"
	"github.com/shieldworks/aegis/core/env"
	"github.com/shieldworks/aegis/core/log"
	"github.com/shieldworks/aegis/core/validation"
	"io"
	"net/http"
	"strings"
)

func List(cid string, w http.ResponseWriter, r *http.Request, svid string) {
	j := audit.JournalEntry{
		CorrelationId: cid,
		Entity:        reqres.SecretListRequest{},
		Method:        r.Method,
		Url:           r.RequestURI,
		Svid:          svid,
		Event:         audit.EventEnter,
	}

	audit.Log(j)

	// Only sentinel can list.
	if !validation.IsSentinel(svid) {
		j.Event = audit.EventBadSvid
		audit.Log(j)

		log.DebugLn(&cid, "List: bad svid", svid)

		w.WriteHeader(http.StatusBadRequest)
		_, err := io.WriteString(w, "")
		if err != nil {
			log.InfoLn(&cid, "List: Problem sending response", err.Error())
		}

		return
	}

	log.TraceLn(&cid, "List: before defer")

	defer func() {
		err := r.Body.Close()
		if err != nil {
			log.InfoLn(&cid, "List: Problem closing body")
		}
	}()

	log.TraceLn(&cid, "List: after defer")

	tmp := strings.Replace(svid, env.SentinelSvidPrefix(), "", 1)
	parts := strings.Split(tmp, "/")
	if len(parts) == 0 {
		j.Event = audit.EventBadPeerSvid
		audit.Log(j)

		w.WriteHeader(http.StatusBadRequest)
		_, err := io.WriteString(w, "")
		if err != nil {
			log.InfoLn(&cid, "List: Problem with svid", svid)
		}
		return
	}

	workloadId := parts[0]
	secrets := state.AllSecrets(cid)

	log.DebugLn(&cid, "List: will send. workload id:", workloadId)

	// RFC3339 is what Go uses internally when marshaling dates.
	// Choosing it to be consistent.
	sfr := reqres.SecretListResponse{
		Secrets: secrets,
	}

	j.Event = audit.EventOk
	j.Entity = sfr
	audit.Log(j)

	resp, err := json.Marshal(sfr)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err := io.WriteString(w, "List: Problem unmarshaling response")
		if err != nil {
			log.InfoLn(&cid, "List: Problem sending response", err.Error())
		}
		return
	}

	log.DebugLn(&cid, "List: before response")

	_, err = io.WriteString(w, string(resp))
	if err != nil {
		log.InfoLn(&cid, "List: Problem sending response", err.Error())
	}

	log.DebugLn(&cid, "List: after response")
}
