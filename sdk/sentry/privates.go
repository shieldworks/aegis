/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                    <aegis.ist>
 *     .\_/.
 */

package sentry

import (
	"bufio"
	"github.com/pkg/errors"
	"github.com/shieldworks/aegis/core/env"
	"os"
)

func saveData(data string) error {
	path := env.SidecarSecretsPath()

	f, err := os.Create(path)
	if err != nil {
		return errors.Wrap(err, "error saving data")
	}

	w := bufio.NewWriter(f)
	_, err = w.WriteString(data)
	if err != nil {
		return errors.Wrap(err, "error saving data")
	}

	err = w.Flush()
	if err != nil {
		return errors.Wrap(err, "error flushing writer")
	}

	return nil
}

func fetchSecrets() error {
	r, eFetch := Fetch()

	// Aegis Safe was successfully queried, but no secrets found.
	// This means someone has deleted the secret. We cannot let
	// the workload linger with the existing secret, so we remove
	// it from the workload too.
	//
	// If the user wants a more fine-tuned control for this case,
	// that is: if the user wants to keep the existing secret even
	// if it has been deleted from Aegis Safe, then the user should
	// use Aegis SDK directly, instead of using Aegis Sidecar.
	if eFetch == ErrSecretNotFound {
		return saveData("")
	}

	v := r.Data
	if v == "" {
		return nil
	}
	return saveData(v)
}
