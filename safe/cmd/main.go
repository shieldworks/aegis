/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                  <aegis.z2h.dev>
 *     .\_/.
 */

package main

import (
	"context"
	"errors"
	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/go-spiffe/v2/spiffetls/tlsconfig"
	"github.com/spiffe/go-spiffe/v2/workloadapi"
	"io"
	"log"
	"net/http"
	"strings"
)

// TODO: get this from environment.
const socketPath = "unix:///spire-agent-socket/agent.sock"

func main() {
	log.Println("Acquiring identity…")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	source, err := workloadapi.NewX509Source(
		ctx, workloadapi.WithClientOptions(workloadapi.WithAddr(socketPath)),
	)
	if err != nil {
		log.Fatalf("Unable to fetch X.509 Bundle: %v", err)
	}
	defer func(source *workloadapi.X509Source) {
		err := source.Close()
		if err != nil {
			log.Printf("Problem closing SVID Bundle source: %v\n", err)
		}
	}(source)

	svid, err := source.GetX509SVID()
	if err != nil {
		log.Fatalf("Unable to get X.509 SVID from source bundle: %v", err)
	}

	svidId := svid.ID
	validSpiffeId := strings.HasPrefix(
		svidId.String(),
		"spiffe://aegis.z2h.dev/ns/aegis-system/sa/aegis-safe/n/",
	)
	if !validSpiffeId {
		log.Fatalf(
			"Svid check: I don’t know you, and it’s crazy: '%s'", svidId.String(),
		)
	}

	log.Println("Acquired identity.")

	// Set up a `/` resource handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Request received")
		urls := r.URL.String()
		log.Println("request.url '", urls, "' path '", r.URL.Path, "' uri '", r.URL.RequestURI(), "'")

		// TODO: POST /v1/fetch  : sidecar to fetch secret
		// TODO: POST /v1/secret : sentinel to upsert secret

		_, _ = io.WriteString(w, "Success!!!")
	})

	// TODO: ability to trust these matchers via Env.
	authorizer := tlsconfig.AdaptMatcher(func(id spiffeid.ID) error {
		if strings.HasPrefix(id.String(), "spiffe://aegis.z2h.dev/workload/") {
			return nil
		}
		return errors.New("TLS Config: I don’t know you, and it’s crazy")
	})

	tlsConfig := tlsconfig.MTLSServerConfig(source, source, authorizer)
	server := &http.Server{
		Addr:      ":8443",
		TLSConfig: tlsConfig,
	}

	if err := server.ListenAndServeTLS("", ""); err != nil {
		log.Fatalf("Error on serve: %v", err)
	}
}
