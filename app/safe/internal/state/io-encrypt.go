/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                    <aegis.ist>
 *     .\_/.
 */

package state

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"filippo.io/age"
	"github.com/pkg/errors"
	"github.com/shieldworks/aegis/core/log"
	"io"
)

func encryptToWriterAge(out io.Writer, data string) error {
	_, publicKey, _ := ageKeyTriplet()
	recipient, err := age.ParseX25519Recipient(publicKey)
	if err != nil {
		return errors.Wrap(err, "encryptToWriterAge: failed to parse public key")
	}

	wrappedWriter, err := age.Encrypt(out, recipient)
	if err != nil {
		return errors.Wrap(err, "encryptToWriterAge: failed to create encrypted file")
	}

	defer func() {
		err := wrappedWriter.Close()
		if err != nil {
			id := "AEGIIOCL"
			log.InfoLn(&id, "encryptToWriterAge: problem closing stream", err.Error())
		}
	}()

	if _, err := io.WriteString(wrappedWriter, data); err != nil {
		return errors.Wrap(err, "encryptToWriterAge: failed to write to encrypted file")
	}

	return nil
}

func encryptToWriterAes(out io.Writer, data string) error {
	_, _, aesKey := ageKeyTriplet()
	aesKeyDecoded, err := hex.DecodeString(aesKey)
	if err != nil {
		return errors.Wrap(err, "encryptToWriter: failed to decode AES key")
	}

	block, err := aes.NewCipher(aesKeyDecoded)
	if err != nil {
		return errors.Wrap(err, "encryptToWriter: failed to create AES cipher block")
	}

	// The IV needs to be unique, but not secure. Therefore, it’s common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, aes.BlockSize+len(data))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], []byte(data))

	_, err = out.Write(ciphertext)
	if err != nil {
		return errors.Wrap(err, "encryptToWriter: failed to write to encrypted file")
	}

	return nil
}
