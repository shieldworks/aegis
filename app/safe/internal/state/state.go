/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                    <aegis.ist>
 *     .\_/.
 */

package state

import (
	"encoding/json"
	"github.com/shieldworks/aegis/app/safe/internal/template"
	entity "github.com/shieldworks/aegis/core/entity/data/v1"
	"github.com/shieldworks/aegis/core/env"
	"github.com/shieldworks/aegis/core/log"
	"os"
	"strings"
	"sync"
	"time"
)

// This is where all the secrets are stored.
var secrets sync.Map

const selfName = "aegis-safe"

type AegisInternalCommand struct {
	LogLevel int `json:"logLevel"`
}

var ageKey = ""
var lock sync.Mutex

func SetAgeKey(k string) {
	lock.Lock()
	defer lock.Unlock()
	ageKey = k
}

func evaluate(data string) *AegisInternalCommand {
	var command AegisInternalCommand
	err := json.Unmarshal([]byte(data), &command)
	if err != nil {
		return nil
	}
	return &command
}

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

	go func() {
		for e := range errChan {
			// If the `persist` operation spews out an error, log it.
			log.ErrorLn("handleSecrets: error persisting secret:", e.Error())
		}
	}()

	for {
		// Buffer overflow check.
		if len(secretQueue) == env.SafeSecretBufferSize() {
			log.ErrorLn(
				"handleSecrets: there are too many k8s secrets queued. " +
					"The goroutine will BLOCK until the queue is cleared.",
			)
		}

		// Get a secret to be persisted to the disk.
		secret := <-secretQueue

		log.TraceLn("picked a secret", len(secretQueue))

		// Persist the secret to disk.
		//
		// Each secret is persisted one at a time, with the order they
		// come in.
		//
		// Do not call this function elsewhere.
		// It is meant to be called inside this `handleSecrets` goroutine.
		persist(secret, errChan)

		log.TraceLn("should have persisted the secret.")
	}
}

func handleK8sSecrets() {
	errChan := make(chan error)

	go func() {
		for e := range errChan {
			// If the `persistK8s` operation spews out an error, log it.
			log.ErrorLn("handleK8sSecrets: error persisting secret:", e.Error())
		}
	}()

	for {
		// Buffer overflow check.
		if len(secretQueue) == env.SafeSecretBufferSize() {
			log.ErrorLn(
				"handleK8sSecrets: there are too many k8s secrets queued. " +
					"The goroutine will BLOCK until the queue is cleared.",
			)
		}

		// Get a secret to be persisted to the disk.
		secret := <-k8sSecretQueue

		log.TraceLn("handleK8sSecrets: picked k8s secret")

		// Sync up the secret to etcd as a Kubernetes Secret.
		//
		// Each secret is synced one at a time, with the order they
		// come in.
		//
		// Do not call this function elsewhere.
		// It is meant to be called inside this `handleK8sSecrets` goroutine.
		persistK8s(secret, errChan)

		log.TraceLn("handleK8sSecrets: Should have persisted k8s secret")
	}
}

func init() {
	go handleSecrets()
	go handleK8sSecrets()
}

var secretsPopulated = false
var secretsPopulatedLock = sync.Mutex{}

func populateSecrets() {
	secretsPopulatedLock.Lock()
	defer secretsPopulatedLock.Unlock()

	root := env.SafeDataPath()
	files, err := os.ReadDir(root)
	if err != nil {
		log.InfoLn("populateSecrets problem:", err.Error())
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

func AllSecrets() []entity.Secret {
	var result []entity.Secret

	// Check existing stored secrets files.
	// If Aegis pod is evicted and revived, it will not have knowledge about
	// the secret it has. This loop helps it re-populate its cache.
	if !secretsPopulated {
		populateSecrets()
	}

	// Range over all existing secrets.
	secrets.Range(func(key any, value any) bool {
		v := value.(entity.SecretStored)

		result = append(result, entity.Secret{
			Name:    v.Name,
			Created: entity.JsonTime(v.Created),
			Updated: entity.JsonTime(v.Updated),
		})

		return true
	})

	if result == nil {
		return []entity.Secret{}
	}

	return result
}

func UpsertSecret(secret entity.SecretStored) {
	if secret.Name == selfName {
		cmd := evaluate(secret.Value)
		if cmd != nil {
			newLogLevel := cmd.LogLevel
			log.InfoLn("Setting new level to:", newLogLevel)
			log.SetLevel(log.Level(newLogLevel))
		}
	}

	s, exists := secrets.Load(secret.Name)
	now := time.Now()
	if exists {
		ss := s.(entity.SecretStored)
		secret.Created = ss.Created
	} else {
		secret.Created = now
	}
	secret.Updated = now

	log.InfoLn("UpsertSecret:",
		"created", secret.Created, "updated", secret.Updated, "name", secret.Name,
	)

	if secret.Value == "" {
		currentState.Decrement(secret.Name)
		secrets.Delete(secret.Name)
	} else {
		parsedStr, err := template.Parse(secret)
		if err != nil {
			log.InfoLn("Error parsing secret. Will use fallback value.", err.Error())
		}

		secret.ValueTransformed = parsedStr
		currentState.Increment(secret.Name)
		secrets.Store(secret.Name, secret)
	}

	store := secret.Meta.BackingStore

	switch store {
	case entity.File:
		log.TraceLn("Will push secret. len", len(secretQueue), "cap", cap(secretQueue))
		secretQueue <- secret
		log.TraceLn("Pushed secret. len", len(secretQueue), "cap", cap(secretQueue))
	case entity.Cluster:
		panic("Cluster backing store not implemented yet!")
	}

	useK8sSecrets := secret.Meta.UseKubernetesSecret
	if useK8sSecrets {
		log.TraceLn(
			"will push Kubernetes secret. len", len(k8sSecretQueue),
			"cap", cap(k8sSecretQueue),
		)
		k8sSecretQueue <- secret
		log.TraceLn(
			"pushed Kubernetes secret. len", len(k8sSecretQueue),
			"cap", cap(k8sSecretQueue),
		)
	}
}

func ReadSecret(key string) *entity.SecretStored {
	result, ok := secrets.Load(key)
	if !ok {
		stored := readFromDisk(key)
		if stored == nil {
			return nil
		}
		currentState.Increment(stored.Name)
		secrets.Store(stored.Name, *stored)
		secretQueue <- *stored
		return stored
	}

	s := result.(entity.SecretStored)
	return &s
}

type StateStatus struct {
	SecretQueueLen int
	SecretQueueCap int
	K8sQueueLen    int
	K8sQueueCap    int
	NumSecrets     int
	lock           *sync.Mutex
}

var currentState = StateStatus{
	SecretQueueLen: 0,
	SecretQueueCap: 0,
	K8sQueueLen:    0,
	K8sQueueCap:    0,
	NumSecrets:     0,
	lock:           &sync.Mutex{},
}

func (s *StateStatus) Increment(name string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	_, ok := secrets.Load(name)
	if !ok {
		s.NumSecrets++
	}
}

func (s *StateStatus) Decrement(name string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	_, ok := secrets.Load(name)
	if ok {
		s.NumSecrets--
	}
}

func Stats() StateStatus {
	currentState.lock.Lock()
	defer currentState.lock.Unlock()

	currentState.K8sQueueCap = cap(k8sSecretQueue)
	currentState.K8sQueueLen = len(k8sSecretQueue)
	currentState.SecretQueueCap = cap(secretQueue)
	currentState.SecretQueueLen = len(secretQueue)

	return currentState
}
