/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                  <aegis.z2h.dev>
 *     .\_/.
 */

package bootstrap

import (
	"aegis-safe/internal/entity/reqres"
	service "aegis-safe/internal/service/v1"
	"aegis-safe/internal/state"
	"context"
	"github.com/go-kit/kit/endpoint"
	"os"
)

func MakeBootstrapEndpoint(svc service.ApiV1Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		r, ok := request.(reqres.BootstrapRequest)
		if !ok {
			return reqres.BootstrapResponse{
				Err: "malformed payload",
			}, nil
		}

		// Already bootstrapped; discard request.
		if state.Bootstrapped() {
			return reqres.BootstrapResponse{}, nil
		}

		// TODO: sanitize these.
		notaryId := r.NotaryId
		workloadToken := r.WorkloadToken
		adminToken := r.AdminToken
		// TODO: empty check
		envNotaryId := os.Getenv("AEGIS_NOTARY_ID")

		if envNotaryId != notaryId {
			return reqres.BootstrapResponse{
				Err: "I don’t know you, and it’s crazy…",
			}, nil
		}

		err := svc.Bootstrap(ctx, adminToken, workloadToken)
		if err != nil {
			return reqres.BootstrapResponse{
				Err: "Unknown problem while bootstrapping Safe.",
			}, nil
		}

		return reqres.BootstrapResponse{}, nil
	}
}
