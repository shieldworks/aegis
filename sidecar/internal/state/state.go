/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                  <aegis.z2h.dev>
 *     .\_/.
 */

package state

import "sync"

var mux = sync.Mutex{}
var id = ""
var secret = ""
var safeApiRoot = ""

func Bootstrap(newId, newSecret, newSafeApiRoot string) {
	mux.Lock()
	defer mux.Unlock()
	id = newId
	secret = newSecret
	safeApiRoot = newSafeApiRoot
}

func Id() string {
	mux.Lock()
	defer mux.Unlock()
	return id
}

func Secret() string {
	mux.Lock()
	defer mux.Unlock()
	return secret
}

func SafeApiRoot() string {
	mux.Lock()
	defer mux.Unlock()
	return safeApiRoot
}
