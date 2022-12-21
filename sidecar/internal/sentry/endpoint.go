/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                  <aegis.z2h.dev>
 *     .\_/.
 */

package sentry

import (
	reqres "aegis-sidecar/internal/entity/reqres/v1"
	"aegis-sidecar/internal/state"
	"bytes"
	"context"
	"encoding/json"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/transport/http"
	"io"
	nhttp "net/http"
	"net/url"
)

func encodeRequest(_ context.Context, r *nhttp.Request, request interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return err
	}
	r.Body = io.NopCloser(&buf)
	return nil
}

func decodeSecretFetchResponse(_ context.Context, r *nhttp.Response) (interface{}, error) {
	var response reqres.SecretFetchResponse
	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}

func newSafeFetchEndpoint() endpoint.Endpoint {
	uu, err := url.JoinPath(state.SafeApiRoot(), "/v1/fetch")
	if err != nil {
		// TODO: handle me
		panic("handle me")
	}
	u, err := url.Parse(uu)
	if err != nil {
		// TODO: handle me
		panic("handle me")
	}
	return http.NewClient(
		"POST", u, encodeRequest, decodeSecretFetchResponse,
	).Endpoint()
}
