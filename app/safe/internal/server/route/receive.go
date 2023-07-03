/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                    <aegis.ist>
 *     .\_/.
 */

package route

import (
	"github.com/shieldworks/aegis/app/safe/internal/state"
	"github.com/shieldworks/aegis/core/audit"
	"github.com/shieldworks/aegis/core/env"
	"github.com/shieldworks/aegis/core/log"
	"net/http"
)

func ReceiveKeys(cid string, w http.ResponseWriter, r *http.Request, svid string) {
	if env.SafeManualKeyInput() && !state.MasterKeySet() {
		log.InfoLn(&cid, "Receive: Master key not set")
		return
	}

	j := createDefaultJournalEntry(cid, svid, r)
	audit.Log(j)

	if !isSentinel(j, cid, w, svid) {
		j.Event = audit.EventBadSvid
		audit.Log(j)
		return
	}

	log.DebugLn(&cid, "ReceiveKeys: sentinel svid:", svid)

	body := readBody(cid, r, w, j)
	if body == nil {
		j.Event = audit.EventBadPayload
		audit.Log(j)
		return
	}

	ur := unmarshalKeyInputRequest(cid, body, j, w)
	if ur == nil {
		j.Event = audit.EventBadPayload
		audit.Log(j)
		return
	}

	sr := *ur

	j.Entity = sr

	aesCipherKey := sr.AesCipherKey
	agePrivateKey := sr.AgeSecretKey
	agePublicKey := sr.AgePublicKey

	keysCombined := agePrivateKey + "\n" + agePublicKey + "\n" + aesCipherKey

	state.SetMasterKey(keysCombined)
}
