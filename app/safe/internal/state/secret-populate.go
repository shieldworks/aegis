/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                    <aegis.ist>
 *     .\_/.
 */

package state

import (
	"github.com/pkg/errors"
	"github.com/shieldworks/aegis/core/env"
	"github.com/shieldworks/aegis/core/log"
	"os"
	"strings"
	"sync"
)

// This is where all the secrets are stored.
var secrets sync.Map

var secretsPopulated = false
var secretsPopulatedLock = sync.Mutex{}

func populateSecrets(cid string) error {
	log.TraceLn(&cid, "populateSecrets: populating secrets...")
	secretsPopulatedLock.Lock()
	defer secretsPopulatedLock.Unlock()

	root := env.SafeDataPath()
	files, err := os.ReadDir(root)
	if err != nil {
		return errors.Wrap(err, "populateSecrets: problem reading secrets directory")
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

		secretOnDisk, err := readFromDisk(key)
		if err != nil {
			log.ErrorLn(&cid, "populateSecrets: problem reading secret from disk:", err.Error())
			continue
		}
		if secretOnDisk != nil {
			currentState.Increment(key)
			secrets.Store(key, *secretOnDisk)
		}
	}

	secretsPopulated = true
	log.TraceLn(&cid, "populateSecrets: secrets populated.")
	return nil
}
