/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                  <aegis.z2h.dev>
 *     .\_/.
 */

package sentry

import (
	v1 "aegis-sidecar/internal/entity/reqres/v1"
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/go-spiffe/v2/spiffetls/tlsconfig"
	"github.com/spiffe/go-spiffe/v2/workloadapi"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func saveData(data string) {
	path := "/opt/aegis/secrets.json"

	f, err := os.Create(path)
	if err != nil {
		// TODO: handle me.
		panic("poop!")
	}

	w := bufio.NewWriter(f)
	_, err = w.WriteString(data)
	if err != nil {
		// TODO: handle me
		panic("poop!")
	}

	err = w.Flush()
	if err != nil {
		// TODO: handle
		panic("poop")
	}

	log.Println("saved secret:", path)
}

// TODO: get this from environment.
const (
	socketPath = "unix:///spire-agent-socket/agent.sock"
	serverUrl  = "https://aegis-safe.aegis-system.svc.cluster.local:8443/"
)

func fetchSecrets() {
	log.Println("fetching secrets…")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	source, err := workloadapi.NewX509Source(
		ctx, workloadapi.WithClientOptions(workloadapi.WithAddr(socketPath)),
	)

	if err != nil {
		log.Println("Failed getting SVID Bundle from the SPIRE Workload API. Will retry.")
		return
	}

	svid, err := source.GetX509SVID()
	if err != nil {
		log.Println("Malformed SVID. will try again.")
		return
	}

	defer func(source *workloadapi.X509Source) {
		if source == nil {
			return
		}
		err := source.Close()
		if err != nil {
			log.Println("Problem closing the workload source.")
		}
	}(source)

	// TODO: add this to technical specs markdown too.
	//
	// SPIFFE ID format:
	//   spiffe://aegis.z2h.dev/workload/$workloadName/ns/{{ .PodMeta.Namespace }}
	//   /sa/{{ .PodSpec.ServiceAccountName }}/n/{{ .PodMeta.Name }}
	//
	// For aegis-system components $workloadName is:
	// - aegis-safe
	// - or aegis-system.
	//
	// For the non-aegis-system workloads that `safe` injects secrets,
	// $workloadName is determined by the workload's ClusterSPIFFEID CRD.

	// Make sure that we are calling Safe from a workload that Aegis knows about.
	if !strings.HasPrefix(
		svid.ID.String(),
		"spiffe://aegis.z2h.dev/workload/",
	) {
		log.Fatalf("Untrusted workload. Killing the container.")
		return
	}

	authorizer := tlsconfig.AdaptMatcher(func(id spiffeid.ID) error {
		if strings.HasPrefix(
			// Only `aegis-safe` can respond to this binary.
			id.String(),
			"spiffe://aegis.z2h.dev/workload/aegis-safe/ns/aegis-system/sa/aegis-safe/n/",
		) {
			return nil
		}

		return errors.New("I don’t know you, and it’s crazy: '" + id.String() + "'")
	})

	p, err := url.JoinPath(serverUrl, "/v1/fetch")
	if err != nil {
		log.Fatalf("Problem generating server url. Killing the container.")
		return
	}

	tlsConfig := tlsconfig.MTLSClientConfig(source, source, authorizer)
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
	}

	sr := v1.SecretFetchRequest{}

	md, err := json.Marshal(sr)
	if err != nil {
		log.Println("Trouble generating payload. Will try again.")
		return
	}

	r, err := client.Post(p, "application/json", bytes.NewBuffer(md))
	if err != nil {
		log.Println("Problem connecting to Aegis Safe API endpoint URL. Will retry")
		return
	}

	defer func(b io.ReadCloser) {
		if b == nil {
			return
		}
		err := b.Close()
		if err != nil {
			log.Println("Problem closing response body.")
		}
	}(r.Body)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println("Unable to read the response body from Aegis Safe. Will retry.")
		return
	}

	var sfr v1.SecretFetchResponse

	err = json.Unmarshal(body, &sfr)
	if err != nil {
		log.Println("Unable to deserialize the response body from Aegis Safe. Will retry.")
		return
	}

	data := sfr.Data
	saveData(data)
}
