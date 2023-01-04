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
	"net/url"
	"strings"
)

// TODO: get this from environment.
const (
	socketPath = "unix:///spire-agent-socket/agent.sock"
	serverUrl  = "https://aegis-safe.aegis-system.svc.cluster.local:8443/"
)

func main() {
	// TODO:
	//
	// 1. Fetch workload SVID + bundle from SPIRE
	// 2. Get SafeSpiffeId from environment.
	// 3. Send a GET request to Safe (safe can know who you are from your spiffeid)
	// 4. parse and safe the returned data.

	log.Println("Welcome to sentinel")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	log.Println("will create svid")

	source, err := workloadapi.NewX509Source(
		ctx, workloadapi.WithClientOptions(workloadapi.WithAddr(socketPath)),
	)

	if err != nil {
		log.Println("Unable to create X509 source")
	} else {
		svid, err := source.GetX509SVID()
		if err != nil {
			// 2023/01/03 19:37:58 svid.id spiffe://aegis.z2h.dev/ns/default/sa/default/n/aegis-workload-demo-559877fd7d-92rcn
			log.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!! could not get svid")
		}
		log.Println("svid.id", svid.ID)

		log.Println("Everything is awesome!", source)
	}
	defer func(source *workloadapi.X509Source) {
		err := source.Close()
		if err != nil {
			// TODO: handle me
		}
	}(source)

	// Allowed SPIFFE ID
	// serverID := spiffeid.RequireFromString("spiffe://example.org/server")
	// spiffe://aegis.z2h.dev/ns/{{ .PodMeta.Namespace }}/sa/{{ .PodSpec.ServiceAccountName }}/n/{{ .PodMeta.Name }}

	authorizer := tlsconfig.AdaptMatcher(func(id spiffeid.ID) error {
		ids := id.String()

		if strings.HasPrefix(ids, "spiffe://aegis.z2h.dev/ns/aegis-system/sa/aegis-safe/n/") {
			return nil
		}

		return errors.New("I don’t know you, and it’s crazy")
	})

	// Create a `tls.Config` to allow mTLS connections, and verify that presented certificate has SPIFFE ID `spiffe://example.org/server`
	tlsConfig := tlsconfig.MTLSClientConfig(source, source, authorizer)
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
	}

	p, err := url.JoinPath(serverUrl, "/v1/fetch")
	if err != nil {
		// TODO: handle this
		return
	}

	r, err := client.Get(p)
	if err != nil {
		log.Fatalf("Error connecting to %q: %v", serverUrl, err)
	}

	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalf("Unable to read body: %v", err)
	}

	log.Printf("%s", body)
}
