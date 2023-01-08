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
	"aegis-sidecar/internal/env"
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/go-spiffe/v2/spiffetls/tlsconfig"
	"github.com/spiffe/go-spiffe/v2/workloadapi"
	"github.com/zerotohero-dev/aegis-core/validation"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

func saveData(data string) {
	path := env.SidecarSecretsPath()

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

func fetchSecrets() {
	log.Println("fetching secrets…")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	source, err := workloadapi.NewX509Source(
		ctx, workloadapi.WithClientOptions(workloadapi.WithAddr(env.SpiffeSocketUrl())),
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

	// Make sure that we are calling Safe from a workload that Aegis knows about.
	if !validation.IsWorkload(svid.ID.String()) {
		log.Fatalf("Untrusted workload. Killing the container.")
		return
	}

	authorizer := tlsconfig.AdaptMatcher(func(id spiffeid.ID) error {
		if validation.IsSafe(id.String()) {
			return nil
		}

		return errors.New("I don’t know you, and it’s crazy: '" + id.String() + "'")
	})

	p, err := url.JoinPath(env.SafeEndpointUrl(), "/v1/fetch")
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
