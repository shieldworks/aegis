/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                    <aegis.ist>
 *     .\_/.
 */

package audit

import (
	reqres "github.com/shieldworks/aegis/core/entity/reqres/safe/v1"
	"github.com/shieldworks/aegis/core/log"
)

type Event string

const EventEnter Event = "aegis-enter"
const EventBadSvid Event = "aegis-bad-svid"
const EventBrokenBody Event = "aegis-broken-body"
const EventRequestTypeMismatch Event = "aegis-request-type-mismatch"
const EventBadPeerSvid Event = "aegis-bad-peer-svid"
const EventNoSecret Event = "aegis-no-secret"
const EventOk Event = "aegis-ok"
const EventNoWorkloadId Event = "aegis-no-workload-id"
const EventNoValue Event = "aegis-no-value"
const EventEncryptionFailed Event = "aegis-encryption-failed"
const EventDecryptionFailed Event = "aegis-decryption-failed"

type JournalEntry struct {
	CorrelationId string
	Entity        any
	Method        string
	Url           string
	Svid          string
	Event         Event
}

func printAudit(correlationId, method, url, svid, message string) {
	log.InfoLn(
		correlationId,
		"method", method,
		"url", url,
		"svid", svid,
		"msg", message,
	)
}

func Log(e JournalEntry) {
	if e.Entity == nil {
		printAudit(
			e.CorrelationId,
			e.Method, e.Url, e.Svid, string(e.Event),
		)
	}

	switch v := e.Entity.(type) {
	case reqres.SecretFetchRequest:
		printAudit(
			e.CorrelationId,
			e.Method, e.Url, e.Svid,
			"e:"+v.Err+"m"+string(e.Event),
		)
	case reqres.SecretFetchResponse:
		printAudit(
			e.CorrelationId,
			e.Method, e.Url, e.Svid,
			"e:"+v.Err+"c:"+v.Created+"u:"+v.Updated+"m:"+string(e.Event),
		)
	case reqres.SecretUpsertRequest:
		printAudit(
			e.CorrelationId,
			e.Method, e.Url, e.Svid,
			"e:"+v.Err+"m"+string(e.Event),
		)
	case reqres.SecretUpsertResponse:
		printAudit(
			e.CorrelationId,
			e.Method, e.Url, e.Svid,
			"e:"+v.Err+"m"+string(e.Event),
		)
	case reqres.SecretListRequest:
		printAudit(
			e.CorrelationId,
			e.Method, e.Url, e.Svid,
			"e:"+v.Err+"m"+string(e.Event),
		)
	case reqres.SecretListResponse:
		printAudit(
			e.CorrelationId,
			e.Method, e.Url, e.Svid,
			"e:"+v.Err+"m"+string(e.Event),
		)
	default:
		printAudit(
			e.CorrelationId,
			e.Method, e.Url, e.Svid,
			"e: UNKNOWN ENTITY in AUDIT LOG",
		)
	}
}
