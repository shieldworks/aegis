/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                    <aegis.ist>
 *     .\_/.
 */

package state

import (
	"bytes"
	"encoding/base64"
	entity "github.com/shieldworks/aegis/core/entity/data/v1"
	"github.com/shieldworks/aegis/core/log"
	"sync"
	"time"
)

const InitialSecretValue = `{"empty":true}`
const BlankAgeKeyValue = "{}"

var ageKey = ""
var lock sync.Mutex

// Initialize starts two goroutines: one to process the secret queue and
// another to process the Kubernetes secret queue. These goroutines are
// responsible for handling queued secrets and persisting them to disk.
func Initialize() {
	go processSecretQueue()
	go processK8sSecretQueue()
	go processSecretDeleteQueue()
	go processK8sSecretDeleteQueue()
}

// SetAgeKey sets the age key to be used for encryption and decryption.
func SetAgeKey(k string) {
	lock.Lock()
	defer lock.Unlock()
	ageKey = k
}

// EncryptValue takes a string value and returns an encrypted and base64-encoded
// representation of the input value. If the encryption process encounters any
// error, it will return an empty string and the corresponding error.
func EncryptValue(value string) (string, error) {
	var out bytes.Buffer

	err := encryptToWriter(&out, value)
	if err != nil {
		return "", err
	}

	base64Str := base64.StdEncoding.EncodeToString(out.Bytes())

	return base64Str, nil
}

// DecryptValue takes a base64-encoded and encrypted string value and returns
// the original, decrypted string. If the decryption process encounters any
// error, it will return an empty string and the corresponding error.
func DecryptValue(value string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(value)
	if err != nil {
		return "", err
	}

	decrypted, err := decryptBytes(decoded)
	if err != nil {
		return "", err
	}

	return string(decrypted), nil
}

// AllSecrets returns a slice of entity.Secret containing all secrets
// currently stored. If no secrets are found, an empty slice is
// returned.
func AllSecrets(cid string) []entity.Secret {
	var result []entity.Secret

	// Check existing stored secrets files.
	// If Aegis pod is evicted and revived, it will not have knowledge about
	// the secret it has. This loop helps it re-populate its cache.
	if !secretsPopulated {
		err := populateSecrets(cid)
		if err != nil {
			log.WarnLn(&cid, "Failed to populate secrets from disk", err.Error())
		}
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

// UpsertSecret takes an entity.SecretStored object and inserts it into
// the in-memory store if it doesn't exist, or updates it if it does. It also
// handles updating the backing store and Kubernetes secrets if necessary.
func UpsertSecret(secret entity.SecretStored) {
	cid := secret.Meta.CorrelationId

	s, exists := secrets.Load(secret.Name)
	now := time.Now()
	if exists {
		ss := s.(entity.SecretStored)
		secret.Created = ss.Created
	} else {
		secret.Created = now
	}
	secret.Updated = now

	log.InfoLn(&cid, "UpsertSecret:",
		"created", secret.Created, "updated", secret.Updated, "name", secret.Name,
	)

	if len(secret.Values) > 0 && secret.Values[0] != "" {
		parsedStr, err := secret.Parse()
		if err != nil {
			log.InfoLn(&cid,
				"UpsertSecret: Error parsing secret. Will use fallback value.", err.Error())
		}

		// TODO: make this plural when `parse` can handle multiple values.
		secret.ValueTransformed = parsedStr
		currentState.Increment(secret.Name)
		secrets.Store(secret.Name, secret)
	}

	store := secret.Meta.BackingStore

	switch store {
	case entity.File:
		log.TraceLn(
			&cid, "UpsertSecret: Will push secret. len", len(secretQueue), "cap", cap(secretQueue))
		secretQueue <- secret
		log.TraceLn(
			&cid, "UpsertSecret: Pushed secret. len", len(secretQueue), "cap", cap(secretQueue))
	}

	useK8sSecrets := secret.Meta.UseKubernetesSecret
	if useK8sSecrets {
		log.TraceLn(
			&cid,
			"UpsertSecret: will push Kubernetes secret. len", len(k8sSecretQueue),
			"cap", cap(k8sSecretQueue),
		)
		k8sSecretQueue <- secret
		log.TraceLn(
			&cid,
			"UpsertSecret: pushed Kubernetes secret. len", len(k8sSecretQueue),
			"cap", cap(k8sSecretQueue),
		)
	}
}

func DeleteSecret(secret entity.SecretStored) {
	cid := secret.Meta.CorrelationId

	s, exists := secrets.Load(secret.Name)
	if !exists {
		log.WarnLn(&cid, "DeleteSecret: Secret does not exist. Cannot delete.", secret.Name)

		ss := s.(entity.SecretStored)
		secret.Created = ss.Created

		return
	}

	ss := s.(entity.SecretStored)

	store := ss.Meta.BackingStore

	switch store {
	case entity.File:
		log.TraceLn(
			&cid, "DeleteSecret: Will delete secret. len", len(secretDeleteQueue), "cap", cap(secretDeleteQueue))
		secretDeleteQueue <- secret
		log.TraceLn(
			&cid, "DeleteSecret: Pushed secret to delete. len", len(secretDeleteQueue), "cap", cap(secretDeleteQueue))
	}

	useK8sSecrets := secret.Meta.UseKubernetesSecret
	if useK8sSecrets {
		log.TraceLn(
			&cid,
			"DeleteSecret: will push Kubernetes secret to delete. len", len(k8sSecretDeleteQueue),
			"cap", cap(k8sSecretDeleteQueue),
		)
		k8sSecretDeleteQueue <- secret
		log.TraceLn(
			&cid,
			"DeleteSecret: pushed Kubernetes secret to delete. len", len(k8sSecretDeleteQueue),
			"cap", cap(k8sSecretDeleteQueue),
		)
	}

	// Remove the secret from the memory.
	currentState.Decrement(secret.Name)
	secrets.Delete(secret.Name)
}

// ReadSecret takes a key string and returns a pointer to an entity.SecretStored
// object if the secret exists in the in-memory store. If the secret is not
// found in memory, it attempts to read it from disk, store it in memory, and
// return it. If the secret is not found on disk, it returns nil.
func ReadSecret(key string) (*entity.SecretStored, error) {
	result, ok := secrets.Load(key)
	if !ok {
		stored, err := readFromDisk(key)
		if err != nil {
			return nil, err
		}

		if stored == nil {
			return nil, nil
		}
		currentState.Increment(stored.Name)
		secrets.Store(stored.Name, *stored)
		secretQueue <- *stored
		return stored, nil
	}

	s := result.(entity.SecretStored)
	return &s, nil
}
