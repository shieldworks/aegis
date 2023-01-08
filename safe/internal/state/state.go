/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                  <aegis.z2h.dev>
 *     .\_/.
 */

package state

import (
	"sync"
)

var secrets sync.Map

func UpsertSecret(id, data string) {
	// log.Println("upsert secret", "id", id, "data", data)
	secrets.Store(id, data)
}

func ReadSecret(key string) string {
	result, ok := secrets.Load(key)
	if !ok {
		return ""
	}

	return result.(string)
}
