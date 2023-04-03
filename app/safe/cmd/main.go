/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                    <aegis.ist>
 *     .\_/.
 */

package main

import (
	"context"
	"github.com/shieldworks/aegis/app/safe/internal/bootstrap"
	"github.com/shieldworks/aegis/app/safe/internal/server"
	"github.com/shieldworks/aegis/core/log"
	"github.com/shieldworks/aegis/core/probe"
)

func main() {
	id := "AEGSAFE"

	log.InfoLn(&id, "Acquiring identity…")

	timedOut := make(chan bool, 1)
	// These channels mus complete in a timely manner, otherwise
	// the timeOut will be fired and will crash the app.
	acquiredSvid := make(chan bool, 1)
	updatedSecret := make(chan bool, 1)
	serverStarted := make(chan bool, 1)

	go bootstrap.NotifyTimeout(timedOut)
	go bootstrap.CreateCryptoKey(&id, updatedSecret)
	go bootstrap.Monitor(&id, acquiredSvid, updatedSecret, serverStarted, timedOut)

	// App is alive; however, not yet ready to accept connections.
	go probe.CreateLiveness()

	ctx, cancel := context.WithCancel(
		context.WithValue(context.Background(), "correlationId", &id),
	)

	defer cancel()

	source := bootstrap.AcquireSource(ctx, acquiredSvid)
	defer func() {
		if err := source.Close(); err != nil {
			log.InfoLn(&id, "Problem closing SVID Bundle source: %v\n", err)
		}
	}()

	if err := server.Serve(source, serverStarted); err != nil {
		log.FatalLn(&id, "failed to serve", err.Error())
	}
}
