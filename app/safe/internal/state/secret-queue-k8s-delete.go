/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                    <aegis.ist>
 *     .\_/.
 */

package state

import (
	entity "github.com/shieldworks/aegis/core/entity/data/v1"
	"github.com/shieldworks/aegis/core/env"
	"github.com/shieldworks/aegis/core/log"
)

// The secrets put here are synced with their Kubernetes Secret counterparts.
var k8sSecretDeleteQueue = make(chan entity.SecretStored, env.SafeSecretBufferSize())

func processK8sSecretDeleteQueue() {
	id := "AEGIHK8D"
	log.InfoLn(&id, "processK8sSecretDeleteQueue: <implement:me>")
}
