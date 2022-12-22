/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                  <aegis.z2h.dev>
 *     .\_/.
 */

package upstream

import (
	reqres "aegis-notary/internal/entity/reqres/v1"
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

func decodeSafeBootstrapResponse(_ context.Context, r *nhttp.Response) (interface{}, error) {
	var response reqres.BootstrapResponse
	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}

func NewSafeBootstrapEndpoint(safeBootstrapUrl string) endpoint.Endpoint {
	u, err := url.Parse(safeBootstrapUrl)
	if err != nil {
		// TODO: handle me
		panic("handle me")
	}
	return http.NewClient(
		"POST", u, encodeRequest, decodeSafeBootstrapResponse,
	).Endpoint()
}

func decodeSidecarBootstrapResponse(_ context.Context, r *nhttp.Response) (interface{}, error) {
	var response reqres.HookResponse
	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}

func NewSidecarBootstrapEndpoint(workloadHookUrl string) endpoint.Endpoint {
	u, err := url.Parse(workloadHookUrl)
	if err != nil {
		// TODO: handle me
		panic("handle me")
	}
	return http.NewClient(
		"POST", u, encodeRequest, decodeSidecarBootstrapResponse,
	).Endpoint()
}
