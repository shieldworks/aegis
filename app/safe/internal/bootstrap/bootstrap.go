/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                    <aegis.ist>
 *     .\_/.
 */

package bootstrap

import (
	"context"
	"filippo.io/age"
	"github.com/shieldworks/aegis/app/safe/internal/state"
	"github.com/shieldworks/aegis/core/env"
	"github.com/shieldworks/aegis/core/log"
	"github.com/shieldworks/aegis/core/probe"
	"github.com/shieldworks/aegis/core/validation"
	"github.com/spiffe/go-spiffe/v2/workloadapi"
	"os"
	"time"
)

func NotifyTimeout(timedOut chan<- bool) {
	time.Sleep(env.SafeSvidRetrievalTimeout())
	timedOut <- true
}

func Monitor(
	correlationId *string,
	acquiredSvid <-chan bool,
	updatedSecret <-chan bool,
	serverStarted <-chan bool,
	timedOut <-chan bool,
) {
	counter := 3
	select {
	case <-acquiredSvid:
		log.InfoLn(correlationId, "Acquired identity.")
		counter--
		if counter == 0 {
			log.DebugLn(correlationId, "Creating readiness probe.")
			go probe.CreateReadiness()
		}
	case <-updatedSecret:
		log.InfoLn(correlationId, "Updated age key.")
		counter--
		if counter == 0 {
			log.DebugLn(correlationId, "Creating readiness probe.")
			go probe.CreateReadiness()
		}
	case <-serverStarted:
		log.InfoLn(correlationId, "Server ready.")
		counter--
		if counter == 0 {
			log.DebugLn(correlationId, "Creating readiness probe.")
			go probe.CreateReadiness()
		}
	case <-timedOut:
		log.FatalLn(correlationId, "Failed to acquire an identity in a timely manner.")
	}
}

func AcquireSource(
	ctx context.Context, acquiredSvid chan<- bool,
) *workloadapi.X509Source {
	source, err := workloadapi.NewX509Source(
		ctx, workloadapi.WithClientOptions(
			workloadapi.WithAddr(env.SpiffeSocketUrl()),
		),
	)

	id := ctx.Value("correlationId").(*string)

	if err != nil {
		log.FatalLn(id, "Unable to fetch X.509 Bundle", err.Error())
	}

	if source == nil {
		log.FatalLn(id, "Could not find source")
	}

	svid, err := source.GetX509SVID()
	if err != nil {
		log.FatalLn(id, "Unable to get X.509 SVID from source bundle", err.Error())
	}

	svidId := svid.ID
	if !validation.IsSafe(svid.ID.String()) {
		log.FatalLn(
			id, "Svid check: I don’t know you, and it’s crazy:", svidId.String(),
		)
	}

	acquiredSvid <- true

	return source
}

func CreateCryptoKey(id *string, updatedSecret chan<- bool) {
	keyPath := env.SafeAgeKeyPath()

	if _, err := os.Stat(keyPath); os.IsNotExist(err) {
		log.FatalLn(id, "CreateCryptoKey: Secret key not mounted at", keyPath)
		return
	}

	data, err := os.ReadFile(keyPath)
	if err != nil {
		log.FatalLn(id, "CreateCryptoKey: Error reading file:", err.Error())
		return
	}

	secret := string(data)

	if secret != state.BlankAgeKeyValue {
		log.InfoLn(id, "Secret has been set in the cluster, will reuse it")
		state.SetAgeKey(secret)
		return
	}

	log.InfoLn(id, "Secret has not been set yet. Will compute a secure secret.")

	identity, err := age.GenerateX25519Identity()
	if err != nil {
		log.FatalLn(id, "Failed to generate key pair: %v", err.Error())
	}

	publicKey := identity.Recipient().String()
	privateKey := identity.String()

	log.TraceLn(id, "Public key: %s...\n", identity.Recipient().String()[:4])
	log.TraceLn(id, "Private key: %s...\n", identity.String()[:16])

	persistKeys(privateKey, publicKey)
	updatedSecret <- true
}
