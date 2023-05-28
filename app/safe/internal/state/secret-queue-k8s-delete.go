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
)

// The secrets put here are synced with their Kubernetes Secret counterparts.
var k8sSecretDeleteQueue = make(chan entity.SecretStored, env.SafeK8sSecretDeleteBufferSize())

func processK8sSecretDeleteQueue() {
	// id := "AEGIHK8D"

	// No need to implement this; but we’ll keep the placeholder here, in case
	// we find a need for it in the future.
	//
	// @see https://github.com/shieldworks/aegis/issues/268
}
