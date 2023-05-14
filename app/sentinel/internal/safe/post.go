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
	"io"
	"log"
	"net/http"
	"net/url"
)

func Post(workloadId, secret, namespace, backingStore string, useKubernetes bool,
	template string, format string, encrypt, deleteSecret, appendSecret bool) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	source, proceed := acquireSource(ctx)
	defer func() {
		if source == nil {
			return
		}
		err := source.Close()
		if err != nil {
			log.Println("Problem closing the workload source.")
		}
	}()
	if !proceed {
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

	f := data.Json
	switch data.SecretFormat(format) {
	case data.Json:
		f = data.Json
	case data.Yaml:
		f = data.Yaml
	default:
		f = data.Json
	}

	sr := reqres.SecretUpsertRequest{
		WorkloadId:    workloadId,
		BackingStore:  bs,
		Namespace:     namespace,
		UseKubernetes: useKubernetes,
		Template:      template,
		Format:        f,
		Encrypt:       encrypt,
		AppendValue:   appendSecret,
		Value:         secret,
	}

	md, err := json.Marshal(sr)
	if err != nil {
		fmt.Println("Trouble generating payload.")
		fmt.Println("")
		return
	}

	var r *http.Response
	if deleteSecret {
		req, err := http.NewRequest(http.MethodDelete, p, bytes.NewBuffer(md))
		if err != nil {
			fmt.Println("Post:Delete: Problem connecting to Aegis Safe API endpoint URL.", err.Error())
			fmt.Println("")
			return
		}
		req.Header.Set("Content-Type", "application/json")
		r, err = client.Do(req)
		if err != nil {
			fmt.Println("Post:Delete: Problem connecting to Aegis Safe API endpoint URL.", err.Error())
			fmt.Println("")
			return
		}
	} else {
		r, err = client.Post(p, "application/json", bytes.NewBuffer(md))
		if err != nil {
			fmt.Println("Post: Problem connecting to Aegis Safe API endpoint URL.", err.Error())
			fmt.Println("")
			return
		}
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
