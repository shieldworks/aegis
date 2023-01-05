/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                  <aegis.z2h.dev>
 *     .\_/.
 */

package validation

import (
	"github.com/spiffe/go-spiffe/v2/workloadapi"
	"log"
	"strings"
)

func EnsureSelfSPIFFEID(source *workloadapi.X509Source) {
	if source == nil {
		log.Fatalf("Could not find source")
	}

	svid, err := source.GetX509SVID()
	if err != nil {
		log.Fatalf("Unable to get X.509 SVID from source bundle: %v", err)
	}

	svidId := svid.ID
	validSpiffeId := strings.HasPrefix(
		svidId.String(),
		"spiffe://aegis.z2h.dev/workload/aegis-safe/ns/aegis-system/sa/aegis-safe/n/",
	)
	if !validSpiffeId {
		log.Fatalf(
			"Svid check: I don’t know you, and it’s crazy: '%s'", svidId.String(),
		)
	}
}
