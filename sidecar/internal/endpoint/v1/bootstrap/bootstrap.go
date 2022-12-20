/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                  <aegis.z2h.dev>
 *     .\_/.
 */

package bootstrap

import (
	service "aegis-sidecar/internal/service/v1"
	"context"
	"github.com/go-kit/kit/endpoint"
)

func MakeBootstrapEndpoint(svc service.ApiV1Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		return nil, nil
	}
}
