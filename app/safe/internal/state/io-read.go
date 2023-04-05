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
	"github.com/pkg/errors"
	entity "github.com/shieldworks/aegis/core/entity/data/v1"
)

// readFromDisk returns a pointer to a secret.
// It returns a nil pointer if secret cannot be read.
func readFromDisk(key string) (*entity.SecretStored, error) {
	contents, err := decryptDataFromDisk(key)
	if err != nil {
		return nil, errors.Wrap(err, "readFromDisk: error decrypting file")
	}

	var secret entity.SecretStored
	err = json.Unmarshal(contents, &secret)
	if err != nil {
		return nil, errors.Wrap(err, "readFromDisk: Failed to unmarshal secret")
	}
	return &secret, nil
}
