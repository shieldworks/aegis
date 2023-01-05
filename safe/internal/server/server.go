/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                  <aegis.z2h.dev>
 *     .\_/.
 */

package server

import (
	"aegis-safe/internal/server/handle"
	"errors"
	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/go-spiffe/v2/spiffetls/tlsconfig"
	"github.com/spiffe/go-spiffe/v2/workloadapi"
	"log"
	"net/http"
	"strings"
)

func Serve(source *workloadapi.X509Source) {
	if source == nil {
		log.Fatalf("Got nil source while trying to serve")
	}

	handle.InitializeRoutes()

	// TODO: ability to trust these matchers via Env.
	authorizer := tlsconfig.AdaptMatcher(func(id spiffeid.ID) error {
		if strings.HasPrefix(id.String(), "spiffe://aegis.z2h.dev/workload/") {
			return nil
		}

		if strings.HasPrefix(id.String(), "spiffe://aegis.z2h.dev/ns/aegis-system/sa/aegis-sentinel/n/") {
			return nil
		}

		return errors.New("TLS Config: I don’t know you, and it’s crazy '" + id.String() + "'")
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
