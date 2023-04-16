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
	"fmt"
	"github.com/shieldworks/aegis/app/safe/internal/state"
	"github.com/shieldworks/aegis/core/audit"
	reqres "github.com/shieldworks/aegis/core/entity/reqres/safe/v1"
	"github.com/shieldworks/aegis/core/env"
	"github.com/shieldworks/aegis/core/log"
	"github.com/shieldworks/aegis/core/validation"
	"io"
	"net/http"
	"strings"
	"time"
)

func Fetch(cid string, w http.ResponseWriter, r *http.Request, svid string) {
	j := audit.JournalEntry{
		CorrelationId: cid,
		Entity:        reqres.SecretFetchRequest{},
		Method:        r.Method,
		Url:           r.RequestURI,
		Svid:          svid,
		Event:         audit.EventEnter,
	}

	audit.Log(j)

	// Only workloads can fetch.
	if !validation.IsWorkload(svid) {
		j.Event = audit.EventBadSvid
		audit.Log(j)

		log.DebugLn(&cid, "Fetch: bad svid", svid)

		w.WriteHeader(http.StatusBadRequest)
		_, err := io.WriteString(w, "")
		if err != nil {
			log.InfoLn(&cid, "Fetch: Problem sending response", err.Error())
		}

		return
	}

	log.DebugLn(&cid, "Fetch: sending response")

	defer func() {
		err := r.Body.Close()
		if err != nil {
			log.InfoLn(&cid, "Fetch: Problem closing body")
		}
	}()

	log.DebugLn(&cid, "Fetch: preparing request")

	tmp := strings.Replace(svid, env.WorkloadSvidPrefix(), "", 1)
	parts := strings.Split(tmp, "/")
	if len(parts) == 0 {
		j.Event = audit.EventBadPeerSvid
		audit.Log(j)

		w.WriteHeader(http.StatusBadRequest)
		_, err := io.WriteString(w, "")
		if err != nil {
			log.InfoLn(&cid, "Fetch: Problem with svid", svid)
		}
		return
	}

	workloadId := parts[0]
	secret, err := state.ReadSecret(workloadId)
	if err != nil {
		log.WarnLn(&cid, "Fetch: Problem reading secret", err.Error())
	}

	log.TraceLn(&cid, "Fetch: workloadId", workloadId)

	// If secret does not exist, send an empty response.
	if secret == nil {
		j.Event = audit.EventNoSecret
		audit.Log(j)

		w.WriteHeader(http.StatusNotFound)
		_, err2 := io.WriteString(w, "")
		if err2 != nil {
			log.InfoLn(&cid, "Fetch: Problem sending response", err2.Error())
		}
		return
	}

	log.DebugLn(&cid, "Fetch: will send. workload id:", workloadId)

	value := ""
	if secret.ValueTransformed != "" {
		value = secret.ValueTransformed
	} else {
		if len(secret.Values) == 1 {
			value = secret.Values[0]
		} else {
			jsonData, err2 := json.Marshal(secret.Values)
			if err2 != nil {
				log.WarnLn(&cid, "Fetch: Problem marshaling values", err2.Error())
			} else {
				value = string(jsonData)
			}
		}
	}

	// RFC3339 is what Go uses internally when marshaling dates.
	// Choosing it to be consistent.
	sfr := reqres.SecretFetchResponse{
		Data:    value,
		Created: fmt.Sprintf("\"%s\"", secret.Created.Format(time.RFC3339)),
		Updated: fmt.Sprintf("\"%s\"", secret.Updated.Format(time.RFC3339)),
	}

	j.Event = audit.EventOk
	j.Entity = sfr
	audit.Log(j)

	resp, err := json.Marshal(sfr)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err2 := io.WriteString(w, "Problem unmarshaling response")
		if err2 != nil {
			log.InfoLn(&cid, "Fetch: Problem sending response", err2.Error())
		}
		return
	}

	log.DebugLn(&cid, "Fetch: before response")

	_, err = io.WriteString(w, string(resp))
	if err != nil {
		log.InfoLn(&cid, "Problem sending response", err.Error())
	}

	log.DebugLn(&cid, "Fetch: after response")
}
