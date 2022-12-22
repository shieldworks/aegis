/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                  <aegis.z2h.dev>
 *     .\_/.
 */

package state

import (
	"log"
	"sync"
)

var once, mutex = NewOnce(), NewSemaphore(1)

// Access to `token` needs to be synchronized since multiple API clients
// can “in theory” read and write it concurrently because the /bootstrap
// api can be called concurrently which spawns separate goroutines per
// network connection (as per how http.Serve behaves).
var adminToken = ""
var workloadToken = ""

func Bootstrapped() bool {
	mutex.Lock()
	defer mutex.Unlock()
	return len(adminToken) > 0
}

func NotaryAdminToken() string {
	mutex.Lock()
	defer mutex.Unlock()
	return adminToken
}

func NotaryWorkloadToken() string {
	mutex.Lock()
	defer mutex.Unlock()
	return workloadToken
}

func Bootstrap(newAdminToken, newWorkloadToken string) {
	// Ensure that the token is set only once.
	once.Do(func() {
		adminToken = newAdminToken
		workloadToken = newWorkloadToken

		// TODO: this is temporary, and definitely NOT for production.
		// Encrypt this with an actual Kubernetes secret
		// limit access to that secret.
		// we can store it in the file system and wrap a CLI around that too.
		log.Println("[SAFE:BOOTSTRAP] token:", adminToken)
	})
}

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

var workloads sync.Map

func RegisterWorkload(id, key string) {
	// log.Println("RegisterWorkload", "id", id, "key", key)
	workloads.Store(id, key)
}

func WorkloadKeyFromId(id string) string {
	result, ok := workloads.Load(id)
	if !ok {
		return ""
	}

	return result.(string)
}
