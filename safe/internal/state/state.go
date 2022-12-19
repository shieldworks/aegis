/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                  <aegis.z2h.dev>
 *     .\_/.
 */

package state

import "sync"

// Access to `token` needs to be synchronized since multiple API clients
// can “in theory” read and write it concurrently because the /bootstrap
// api can be called concurrently which spawns separate goroutines per
// network connection (as per how http.Serve behaves).
// Synchronizing access to `token` using a `channel` will turn out to
//
// be much messier than using a mutex, so we’ll lock over `mux` instead.
var token = ""
var mux = sync.Mutex{}

func Bootstrapped() bool {
	mux.Lock()
	defer mux.Unlock()
	return len(token) > 0
}

func NotaryToken() string {
	mux.Lock()
	defer mux.Unlock()
	return token
}

func SetNotaryToken(newToken string) {
	mux.Lock()
	defer mux.Unlock()
	if len(token) > 0 {
		return
	}
	token = newToken
}
