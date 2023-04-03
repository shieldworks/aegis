/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                    <aegis.ist>
 *     .\_/.
 */

package server

import (
	"github.com/pkg/errors"
	"github.com/shieldworks/aegis/app/safe/internal/server/handle"
	"github.com/shieldworks/aegis/core/env"
	"github.com/shieldworks/aegis/core/probe"
	"github.com/shieldworks/aegis/core/validation"
	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/go-spiffe/v2/spiffetls/tlsconfig"
	"github.com/spiffe/go-spiffe/v2/workloadapi"
	"net/http"
)

func Serve(source *workloadapi.X509Source, serverStarted chan<- bool) error {
	if source == nil {
		return errors.New("serve: got nil source while trying to serve")
	}

	handle.InitializeRoutes()

	authorizer := tlsconfig.AdaptMatcher(func(id spiffeid.ID) error {
		if validation.IsWorkload(id.String()) {
			return nil
		}

		return errors.New(
			"TLS Config: I don’t know you, and it’s crazy '" + id.String() + "'",
		)
	})

	tlsConfig := tlsconfig.MTLSServerConfig(source, source, authorizer)
	server := &http.Server{
		Addr:      env.TlsPort(),
		TLSConfig: tlsConfig,
	}

	serverStarted <- true

	// Since server has started, we can enable the readiness probe.
	go probe.CreateReadiness()

	if err := server.ListenAndServeTLS("", ""); err != nil {
		return errors.Wrap(err, "serve: failed to listen and serve")
	}

	return nil
}
