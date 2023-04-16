/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                    <aegis.ist>
 *     .\_/.
 */

package safe

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	data "github.com/shieldworks/aegis/core/entity/data/v1"
	reqres "github.com/shieldworks/aegis/core/entity/reqres/safe/v1"
	"github.com/shieldworks/aegis/core/env"
	"github.com/shieldworks/aegis/core/validation"
	"github.com/spiffe/go-spiffe/v2/spiffeid"
	"github.com/spiffe/go-spiffe/v2/spiffetls/tlsconfig"
	"github.com/spiffe/go-spiffe/v2/workloadapi"
	"io"
	"log"
	"net/http"
	"net/url"
)

func Post(workloadId, secret, namespace, backingStore string, useKubernetes bool,
	template string, format string, encrypt bool) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	source, err := workloadapi.NewX509Source(
		ctx, workloadapi.WithClientOptions(workloadapi.WithAddr(env.SpiffeSocketUrl())),
	)

	if err != nil {
		fmt.Println("Post: I cannot execute command because I cannot talk to SPIRE.")
		fmt.Println("")
		return
	}

	svid, err := source.GetX509SVID()
	if err != nil {
		fmt.Println("Post: I am having trouble fetching my identity from SPIRE.")
		fmt.Println("Post: I won’t proceed until you put me in a secured container.")
		fmt.Println("")
		return
	}

	defer func() {
		err := source.Close()
		if err != nil {
			log.Println("Post: Problem closing the workload source.", err.Error())
		}
	}()

	// Make sure that the binary is enclosed in a Pod that we trust.
	if !validation.IsSentinel(svid.ID.String()) {
		fmt.Println("Post: I don’t know you, and it’s crazy: '" + svid.ID.String() + "'")
		fmt.Println("`aegis` can only run from within the Sentinel container.")
		fmt.Println("")
		return
	}

	authorizer := tlsconfig.AdaptMatcher(func(id spiffeid.ID) error {
		if validation.IsSafe(id.String()) {
			return nil
		}

		return errors.New("Post: I don’t know you, and it’s crazy: '" + id.String() + "'")
	})

	p, err := url.JoinPath(env.SafeEndpointUrl(), "/sentinel/v1/secrets")
	if err != nil {
		fmt.Println("Post: I am having problem generating Aegis Safe secrets api endpoint URL.", err.Error())
		fmt.Println("")
		return
	}

	tlsConfig := tlsconfig.MTLSClientConfig(source, source, authorizer)
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
	}

	bs := env.SafeBackingStore()
	if backingStore != "" {
		b := data.BackingStore(backingStore)
		switch b {
		case data.File:
			bs = data.File
		case data.Memory:
			bs = data.Memory
		default:
			bs = data.Memory
		}
	}

	f := data.None
	switch data.SecretFormat(format) {
	case data.Json:
		f = data.Json
	case data.Yaml:
		f = data.Yaml
	default:
		f = data.None
	}

	sr := reqres.SecretUpsertRequest{
		WorkloadId:    workloadId,
		BackingStore:  bs,
		Namespace:     namespace,
		UseKubernetes: useKubernetes,
		Template:      template,
		Format:        f,
		Encrypt:       encrypt,
		Value:         secret,
	}

	md, err := json.Marshal(sr)
	if err != nil {
		fmt.Println("Trouble generating payload.")
		fmt.Println("")
		return
	}

	r, err := client.Post(p, "application/json", bytes.NewBuffer(md))
	if err != nil {
		fmt.Println("Post: Problem connecting to Aegis Safe API endpoint URL.", err.Error())
		fmt.Println("")
		return
	}

	defer func(b io.ReadCloser) {
		if b == nil {
			return
		}
		err := b.Close()
		if err != nil {
			log.Println("Post: Problem closing request body.", err.Error())
		}
	}(r.Body)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Post: Unable to read the response body from Aegis Safe.", err.Error())
		fmt.Println("")
		return
	}

	fmt.Println("")
	fmt.Println(string(body))
	fmt.Println("")
}
