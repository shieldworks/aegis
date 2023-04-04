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
	"os"
	"strings"
	"sync"
)

// This is where all the secrets are stored.
var secrets sync.Map

// These are persisted to files. They are buffered, so that they can
// be written in the order they are queued and there are no concurrent
// writes to the same file at a time. An alternative approach would be
// to have a map of queues of `SecretsStored`s per file name but that
// feels like an overkill.
var secretQueue = make(chan entity.SecretStored, env.SafeSecretBufferSize())

// The secrets put here are synced with their Kubernetes Secret counterparts.
var k8sSecretQueue = make(chan entity.SecretStored, env.SafeSecretBufferSize())

func handleSecrets() {
	errChan := make(chan error)

	id := "AEGIHSCR"

	go func() {
		for e := range errChan {
			// If the `persist` operation spews out an error, log it.
			log.ErrorLn(&id, "handleSecrets: error persisting secret:", e.Error())
		}
	}()

	for {
		// Buffer overflow check.
		if len(secretQueue) == env.SafeSecretBufferSize() {
			log.ErrorLn(
				&id,
				"handleSecrets: there are too many k8s secrets queued. "+
					"The goroutine will BLOCK until the queue is cleared.",
			)
		}

		// Get a secret to be persisted to the disk.
		secret := <-secretQueue

		cid := secret.Meta.CorrelationId

		log.TraceLn(&cid, "handleSecrets: picked a secret", len(secretQueue))

		// Persist the secret to disk.
		//
		// Each secret is persisted one at a time, with the order they
		// come in.
		//
		// Do not call this function elsewhere.
		// It is meant to be called inside this `handleSecrets` goroutine.
		persist(secret, errChan)

		log.TraceLn(&cid, "handleSecrets: should have persisted the secret.")
	}
}

func handleK8sSecrets() {
	errChan := make(chan error)

	id := "AEGIHK8S"

	go func() {
		for e := range errChan {
			// If the `persistK8s` operation spews out an error, log it.
			log.ErrorLn(&id, "handleK8sSecrets: error persisting secret:", e.Error())
		}
	}()

	for {
		// Buffer overflow check.
		if len(secretQueue) == env.SafeSecretBufferSize() {
			log.ErrorLn(
				&id,
				"handleK8sSecrets: there are too many k8s secrets queued. "+
					"The goroutine will BLOCK until the queue is cleared.",
			)
		}

		// Get a secret to be persisted to the disk.
		secret := <-k8sSecretQueue

		cid := secret.Meta.CorrelationId

		log.TraceLn(&cid, "handleK8sSecrets: picked k8s secret")

		// Sync up the secret to etcd as a Kubernetes Secret.
		//
		// Each secret is synced one at a time, with the order they
		// come in.
		//
		// Do not call this function elsewhere.
		// It is meant to be called inside this `handleK8sSecrets` goroutine.
		persistK8s(secret, errChan)

		log.TraceLn(&cid, "handleK8sSecrets: Should have persisted k8s secret")
	}
}

var secretsPopulated = false
var secretsPopulatedLock = sync.Mutex{}

func populateSecrets() error {
	secretsPopulatedLock.Lock()
	defer secretsPopulatedLock.Unlock()

	root := env.SafeDataPath()
	files, err := os.ReadDir(root)
	if err != nil {
		log.InfoLn("populateSecrets: problem:", err.Error())
		return
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		fn := file.Name()
		if strings.HasSuffix(fn, ".backup") {
			continue
		}

		key := strings.Replace(fn, ".age", "", 1)

		_, exists := secrets.Load(key)
		if exists {
			continue
		}

		secretOnDisk := readFromDisk(key)
		if secretOnDisk != nil {
			currentState.Increment(key)
			secrets.Store(key, *secretOnDisk)
		}
	}

	secretsPopulated = true
	log.InfoLn("populateSecrets: secrets populated.")
}

func init() {
	go handleSecrets()
	go handleK8sSecrets()
}
