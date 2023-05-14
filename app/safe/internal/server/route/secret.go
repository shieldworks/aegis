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
	"io"
	"net/http"
)

func Secret(cid string, w http.ResponseWriter, r *http.Request, svid string) {
	j := audit.JournalEntry{
		CorrelationId: cid,
		Entity:        reqres.SecretFetchRequest{},
		Method:        r.Method,
		Url:           r.RequestURI,
		Svid:          svid,
		Event:         audit.EventEnter,
	}

	audit.Log(j)

	if !isSentinel(j, cid, w, svid) {
		return
	}

	log.DebugLn(&cid, "Secret: sentinel svid:", svid)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		j.Event = audit.EventBrokenBody
		audit.Log(j)

		w.WriteHeader(http.StatusBadRequest)
		_, err := io.WriteString(w, "")
		if err != nil {
			log.InfoLn(&cid, "Secret: Problem sending response", err.Error())
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

	var sr reqres.SecretUpsertRequest
	err = json.Unmarshal(body, &sr)
	if err != nil {
		j.Event = audit.EventRequestTypeMismatch
		audit.Log(j)
		w.WriteHeader(http.StatusBadRequest)
		_, err := io.WriteString(w, "")
		if err != nil {
			log.InfoLn(&cid, "Secret: Problem sending response", err.Error())
		}
		return
	}

	j.Entity = sr

	workloadId := sr.WorkloadId
	value := sr.Value

	backingStore := sr.BackingStore
	useK8s := sr.UseKubernetes
	namespace := sr.Namespace
	template := sr.Template
	format := sr.Format
	encrypt := sr.Encrypt
	appendValue := sr.AppendValue

	if workloadId == "" && encrypt {
		if value == "" {
			j.Event = audit.EventNoValue
			audit.Log(j)

			w.WriteHeader(http.StatusBadRequest)
			_, err := io.WriteString(w, "")
			if err != nil {
				log.InfoLn(&cid, "Secret: Problem sending response", err.Error())
			}
			return
		}

		encrypted, err := state.EncryptValue(value)
		if err != nil {
			j.Event = audit.EventEncryptionFailed
			audit.Log(j)

			w.WriteHeader(http.StatusInternalServerError)
			_, err := io.WriteString(w, "")
			if err != nil {
				log.InfoLn(&cid, "Secret: Problem sending response", err.Error())
			}
			return
		}

		_, err = io.WriteString(w, encrypted)
		if err != nil {
			log.InfoLn(&cid, "Secret: Problem sending response", err.Error())
		}
		return
	}

	if namespace == "" {
		namespace = "default"
	}

	log.DebugLn(&cid, "Secret:Upsert: ",
		"workloadId:", workloadId,
		"namespace:", namespace,
		"backingStore:", backingStore,
		"template:", template,
		"format:", format,
		"encrypt:", encrypt,
		"appendValue:", appendValue,
		"useK8s", useK8s)

	if workloadId == "" && !encrypt {
		j.Event = audit.EventNoWorkloadId
		audit.Log(j)

		return
	}

	// `encrypt` means that the value is encrypted, so we need to decrypt it.
	if encrypt {
		decrypted, err := state.DecryptValue(value)
		if err != nil {
			j.Event = audit.EventDecryptionFailed
			audit.Log(j)

			w.WriteHeader(http.StatusInternalServerError)
			_, err := io.WriteString(w, "")
			if err != nil {
				log.InfoLn(&cid, "Secret: Problem sending response", err.Error())
			}
			return
		}

		value = decrypted
	}

	if len(value) > 65536 {
		panic("This is just a reminder to implement multiple-valued secrets")
	}

	state.UpsertSecret(entity.SecretStored{
		Name: workloadId,
		Meta: entity.SecretMeta{
			UseKubernetesSecret: useK8s,
			Namespace:           namespace,
			BackingStore:        backingStore,
			Template:            template,
			Format:              format,
			CorrelationId:       cid,
		},
		Values: []string{value},
	}, appendValue)
	log.DebugLn(&cid, "Secret:UpsertEnd: workloadId", workloadId)

	j.Event = audit.EventOk
	audit.Log(j)

	_, err = io.WriteString(w, "OK")
	if err != nil {
		log.InfoLn(&cid, "Secret: Problem sending response", err.Error())
	}
}
