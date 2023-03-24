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
	r, _ := Fetch()
	v := r.Data
	if v == "" {
		return nil
	}
	return saveData(v)
}
