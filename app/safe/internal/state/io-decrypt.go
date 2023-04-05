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
	"filippo.io/age"
	"github.com/pkg/errors"
	"github.com/shieldworks/aegis/core/env"
	"io"
	"os"
	"path"
)

func decryptBytes(data []byte) ([]byte, error) {
	privateKey, _ := ageKeyPair()

	identity, err := age.ParseX25519Identity(privateKey)
	if err != nil {
		return []byte{}, errors.Wrap(err, "decryptBytes: failed to parse private key")
	}

	if len(data) == 0 {
		return []byte{}, errors.Wrap(err, "decryptBytes: file on disk appears to be empty")
	}

	out := &bytes.Buffer{}
	f := bytes.NewReader(data)

	r, err := age.Decrypt(f, identity)
	if err != nil {
		return []byte{}, errors.Wrap(err, "decryptBytes: failed to open encrypted file")
	}

	if _, err := io.Copy(out, r); err != nil {
		return []byte{}, errors.Wrap(err, "decryptBytes: failed to read encrypted file")
	}

	return out.Bytes(), nil
}

func decryptDataFromDisk(key string) ([]byte, error) {
	dataPath := path.Join(env.SafeDataPath(), key+".age")

	if _, err := os.Stat(dataPath); os.IsNotExist(err) {
		return nil, errors.Wrap(err, "decryptDataFromDisk: No file at: "+dataPath)
	}

	data, err := os.ReadFile(dataPath)
	if err != nil {
		return nil, errors.Wrap(err, "decryptDataFromDisk: Error reading file")
	}

	return decryptBytes(data)
}
