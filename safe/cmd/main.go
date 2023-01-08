/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                  <aegis.z2h.dev>
 *     .\_/.
 */

package main

import (
	"aegis-safe/internal/env"
	"aegis-safe/internal/server"
	"aegis-safe/internal/validation"
	"context"
	"github.com/spiffe/go-spiffe/v2/workloadapi"
	"log"
)

func main() {
	log.Println("Acquiring identity…")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	source, err := workloadapi.NewX509Source(
		ctx, workloadapi.WithClientOptions(workloadapi.WithAddr(env.SpiffeSocketUrl())),
	)

	if err != nil {
		log.Fatalf("Unable to fetch X.509 Bundle: %v", err)
	}

	defer func(source *workloadapi.X509Source) {
		if source == nil {
			return
		}
		err := source.Close()
		if err != nil {
			log.Printf("Problem closing SVID Bundle source: %v\n", err)
		}
	}(source)

	validation.EnsureSelfSPIFFEID(source)
	log.Println("Acquired identity.")

	server.Serve(source)
}
