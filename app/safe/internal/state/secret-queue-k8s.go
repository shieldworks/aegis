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
var k8sSecretQueue = make(chan entity.SecretStored, env.SafeSecretBufferSize())

func processK8sSecretQueue() {
	errChan := make(chan error)

	id := "AEGIHK8S"

	go func() {
		for e := range errChan {
			// If the `persistK8s` operation spews out an error, log it.
			log.ErrorLn(&id, "processK8sSecretQueue: error persisting secret:", e.Error())
		}
	}()

	for {
		// Buffer overflow check.
		if len(secretQueue) == env.SafeSecretBufferSize() {
			log.ErrorLn(
				&id,
				"processK8sSecretQueue: there are too many k8s secrets queued. "+
					"The goroutine will BLOCK until the queue is cleared.",
			)
		}

		// Get a secret to be persisted to the disk.
		secret := <-k8sSecretQueue

		cid := secret.Meta.CorrelationId

		log.TraceLn(&cid, "processK8sSecretQueue: picked k8s secret")

		// Sync up the secret to etcd as a Kubernetes Secret.
		//
		// Each secret is synced one at a time, with the order they
		// come in.
		//
		// Do not call this function elsewhere.
		// It is meant to be called inside this `processK8sSecretQueue` goroutine.
		persistK8s(secret, errChan)

		log.TraceLn(&cid, "processK8sSecretQueue: Should have persisted k8s secret")
	}
}
