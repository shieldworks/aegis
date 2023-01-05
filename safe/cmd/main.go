/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                  <aegis.z2h.dev>
 *     .\_/.
 */

package main

import (
	v1 "aegis-safe/internal/entity/reqres/v1"
	"aegis-safe/internal/state"
	"context"
	"encoding/json"
	"errors"
	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/go-spiffe/v2/spiffetls/tlsconfig"
	"github.com/spiffe/go-spiffe/v2/svid/x509svid"
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
		"spiffe://aegis.z2h.dev/workload/aegis-safe/ns/aegis-system/sa/aegis-safe/n/",
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

		if r == nil {
			return
		}

		tlsConnectionState := r.TLS

		if len(tlsConnectionState.PeerCertificates) == 0 {
			log.Println("no peer certs :(")
			return
		}

		id, err := x509svid.IDFromCert(tlsConnectionState.PeerCertificates[0])
		if err != nil {
			log.Println("poop!")
			return
		}

		log.Println("GOT svid:", id.String())

		p := r.URL.Path

		// sidecar -> safe : fetch secrets
		if r.Method == http.MethodPost && p == "/v1/fetch" {
			// TODO: svid validation too.

			body, err := io.ReadAll(r.Body)
			if err != nil {
				// TODO: handle me.
			}
			defer func(Body io.ReadCloser) {
				err := Body.Close()
				if err != nil {
					// TODO: handle me.
				}
			}(r.Body)

			var sr v1.SecretFetchRequest

			err = json.Unmarshal(body, &sr)
			if err != nil {
				// TODO: handle me.
				log.Println("error unmarshalling json")
				return
			}

			// TODO: this shall be parsed from svid instead.
			workloadId := "aegis-workload-demo"

			value := state.ReadSecret(workloadId)

			log.Println("upsert: key: '", workloadId, "' value: '", value, "'")
			log.Println("read:", state.ReadSecret(workloadId))

			sfr := v1.SecretFetchResponse{
				Data: value,
			}

			resp, err := json.Marshal(sfr)
			if err != nil {
				// TODO: handle me
				log.Println("poop!")
			}

			_, _ = io.WriteString(w, string(resp))

			return
		}

		// sentinel -> safe : put secrets
		if r.Method == http.MethodPost && p == "/v1/secret" {
			// TODO: svid validation too.

			body, err := io.ReadAll(r.Body)
			if err != nil {
				// TODO: handle me.
			}
			defer func(Body io.ReadCloser) {
				err := Body.Close()
				if err != nil {
					// TODO: handle me.
				}
			}(r.Body)

			var sr v1.SecretUpsertRequest

			err = json.Unmarshal(body, &sr)
			if err != nil {
				// TODO: handle me.
				log.Println("error unmarshalling json")
				return
			}

			workloadId := sr.WorkloadId
			value := sr.Value

			state.UpsertSecret(workloadId, value)

			log.Println("upsert: key: '", workloadId, "' value: '", value, "'")
			log.Println("read:", state.ReadSecret(workloadId))

			// 2023/01/04 18:33:59 GOT svid: spiffe://aegis.z2h.dev/ns/aegis-system/sa/aegis-sentinel/n/aegis-sentinel-66d445698d-x6m7m
			// 2023/01/04 18:33:59 body ' {"key":"aegis-workload-demo","value":"{\"u\": \"root\", \"p\": \"toppyTopSecret\", \"realm\": \"narnia\"}"} '

			log.Println("body '", string(body), "'")

			_, _ = io.WriteString(w, "OK")
			return
		}

		// TODO: return an error instead.
		_, _ = io.WriteString(w, "OK")
	})

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

	log.Println("Before creating tls config")

	tlsConfig := tlsconfig.MTLSServerConfig(source, source, authorizer)
	server := &http.Server{
		Addr:      ":8443",
		TLSConfig: tlsConfig,
	}

	log.Println("Created tls config")

	if err := server.ListenAndServeTLS("", ""); err != nil {
		log.Fatalf("Error on serve: %v", err)
	}

	// TODO: Probably never executes:
	log.Println("Server started.")
}
