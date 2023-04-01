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
	"github.com/shieldworks/aegis/app/safe/internal/template"
	entity "github.com/shieldworks/aegis/core/entity/data/v1"
	"github.com/shieldworks/aegis/core/log"
	"sync"
	"time"
)

var ageKey = ""
var lock sync.Mutex

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

// UpsertSecret takes an entity.SecretStored object and inserts it into
// the in-memory store if it doesn't exist, or updates it if it does. It also
// handles updating the backing store and Kubernetes secrets if necessary.
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

// ReadSecret takes a key string and returns a pointer to an entity.SecretStored
// object if the secret exists in the in-memory store. If the secret is not
// found in memory, it attempts to read it from disk, store it in memory, and
// return it. If the secret is not found on disk, it returns nil.
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
