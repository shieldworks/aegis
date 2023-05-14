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
	entity "github.com/shieldworks/aegis/core/entity/data/v1"
	reqres "github.com/shieldworks/aegis/core/entity/reqres/safe/v1"
	"github.com/shieldworks/aegis/core/log"
	"github.com/shieldworks/aegis/core/validation"
	"io"
	"net/http"
)

func isSentinel(j audit.JournalEntry, cid string, w http.ResponseWriter, svid string) bool {
	audit.Log(j)

	if validation.IsSentinel(svid) {
		return true
	}

	j.Event = audit.EventBadSvid
	audit.Log(j)

	w.WriteHeader(http.StatusBadRequest)
	_, err := io.WriteString(w, "")
	if err != nil {
		log.InfoLn(&cid, "Delete: Problem sending response", err.Error())
	}

	return false
}

func Delete(cid string, w http.ResponseWriter, r *http.Request, svid string) {
	j := audit.JournalEntry{
		CorrelationId: cid,
		Entity:        reqres.SecretDeleteRequest{},
		Method:        r.Method,
		Url:           r.RequestURI,
		Svid:          svid,
		Event:         audit.EventEnter,
	}

	if !isSentinel(j, cid, w, svid) {
		return
	}

	log.DebugLn(&cid, "Delete: sentinel svid:", svid)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		j.Event = audit.EventBrokenBody
		audit.Log(j)

		w.WriteHeader(http.StatusBadRequest)
		_, err := io.WriteString(w, "")
		if err != nil {
			log.InfoLn(&cid, "Delete: Problem sending response", err.Error())
		}
		return
	}
	defer func(b io.ReadCloser) {
		if b == nil {
			return
		}
		err := b.Close()
		if err != nil {
			log.InfoLn(&cid, "Secret: Problem closing body", err.Error())
		}
	}(r.Body)

	log.DebugLn(&cid, "Secret: Parsed request body")

	var sr reqres.SecretDeleteRequest
	err = json.Unmarshal(body, &sr)
	if err != nil {
		j.Event = audit.EventRequestTypeMismatch
		audit.Log(j)
		w.WriteHeader(http.StatusBadRequest)
		_, err := io.WriteString(w, "")
		if err != nil {
			log.InfoLn(&cid, "Delete: Problem sending response", err.Error())
		}
		return
	}

	j.Entity = sr

	workloadId := sr.WorkloadId

	if workloadId == "" {
		j.Event = audit.EventNoWorkloadId
		audit.Log(j)
		return
	}

	log.DebugLn(&cid, "Secret:Delete: ", "workloadId:", workloadId)

	if workloadId == "" {
		j.Event = audit.EventNoWorkloadId
		audit.Log(j)

		return
	}

	state.DeleteSecret(entity.SecretStored{
		Name: workloadId,
		Meta: entity.SecretMeta{
			CorrelationId: cid,
		},
	})
	log.DebugLn(&cid, "Delete:End: workloadId", workloadId)

	j.Event = audit.EventOk
	audit.Log(j)

	_, err = io.WriteString(w, "OK")
	if err != nil {
		log.InfoLn(&cid, "Delete: Problem sending response", err.Error())
	}
}
