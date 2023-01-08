/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                  <aegis.z2h.dev>
 *     .\_/.
 */

package main

import (
	"aegis-sentinel/internal/env"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/akamensky/argparse"
	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/go-spiffe/v2/spiffetls/tlsconfig"
	"github.com/spiffe/go-spiffe/v2/workloadapi"
	"github.com/zerotohero-dev/aegis/core/entity/reqres/v1"
	"github.com/zerotohero-dev/aegis/core/validation"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

func main() {
	parser := argparse.NewParser(
		"aegis",
		"Assigns secrets to workloads.",
	)

	workload := parser.String(
		"w", "workload",
		&argparse.Options{
			Required: true,
			Help:     "name of the workload (i.e. the '$name' segment of its ClusterSPIFFEID ('spiffe://trustDomain/workload/$name/…'))",
		},
	)

	secret := parser.String(
		"s", "secret",
		&argparse.Options{
			Required: true,
			Help:     "the secret to store for the workload",
		},
	)

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
	}

	if workload == nil || *workload == "" {
		fmt.Println("Please provide a workload name.")
		fmt.Println("")
		fmt.Println("type `aegis -h` (without backticks) and press return for help.")
		fmt.Println("")
		return
	}

	if secret == nil || *secret == "" {
		fmt.Println("Please provide a secret.")
		fmt.Println("")
		fmt.Println("type `aegis -h` (without backticks) and press return for help.")
		fmt.Println("")
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	source, err := workloadapi.NewX509Source(
		ctx, workloadapi.WithClientOptions(workloadapi.WithAddr(env.SpiffeSocketUrl())),
	)

	if err != nil {
		fmt.Println("I cannot execute command because I cannot talk to SPIRE.")
		fmt.Println("")
		return
	}

	svid, err := source.GetX509SVID()
	if err != nil {
		fmt.Println("I am having trouble fetching your identity from SPIRE.")
		fmt.Println("I won’t proceed until you put me in a secured container.")
		fmt.Println("")
		return
	}

	defer func(source *workloadapi.X509Source) {
		err := source.Close()
		if err != nil {
			log.Println("Problem closing the workload source.")
		}
	}(source)

	// Make sure that the binary is enclosed in a Pod that we trust.
	if !validation.IsSentinel(svid.ID.String()) {
		fmt.Println("I don’t know you, and it’s crazy: '" + svid.ID.String() + "'")
		fmt.Println("`aegis` can only run from within the Sentinel container.")
		fmt.Println("")
		return
	}

	authorizer := tlsconfig.AdaptMatcher(func(id spiffeid.ID) error {
		if validation.IsSafe(id.String()) {
			return nil
		}

		return errors.New("I don’t know you, and it’s crazy: '" + id.String() + "'")
	})

	p, err := url.JoinPath(env.SafeEndpointUrl(), "/v1/secret")
	if err != nil {
		fmt.Println("I am having problem generating Aegis Safe secrets api endpoint URL.")
		fmt.Println("")
		return
	}

	tlsConfig := tlsconfig.MTLSClientConfig(source, source, authorizer)
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
	}

	sr := v1.SecretUpsertRequest{
		WorkloadId: *workload,
		Value:      *secret,
	}

	md, err := json.Marshal(sr)
	if err != nil {
		fmt.Println("Trouble generating payload.")
		fmt.Println("")
		return
	}

	r, err := client.Post(p, "application/json", bytes.NewBuffer(md))
	if err != nil {
		fmt.Println("Problem connecting to Aegis Safe API endpoint URL.")
		fmt.Println("")
		return
	}

	defer func(b io.ReadCloser) {
		if b == nil {
			return
		}
		err := b.Close()
		if err != nil {
			log.Println("Problem closing request body.")
		}
	}(r.Body)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Unable to read the response body from Aegis Safe.")
		fmt.Println("")
		return
	}

	fmt.Println("")
	fmt.Println(string(body))
	fmt.Println("")
}
