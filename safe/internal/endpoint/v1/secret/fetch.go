/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                  <aegis.z2h.dev>
 *     .\_/.
 */

package secret

import (
	reqres "aegis-safe/internal/entity/reqres/v1"
	service "aegis-safe/internal/service/v1"
	"aegis-safe/internal/state"
	"context"
	"github.com/go-kit/kit/endpoint"
)

func MakeSecretFetchEndpoint(svc service.ApiV1Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		r, ok := request.(reqres.SecretFetchRequest)
		if !ok {
			return reqres.SecretFetchResponse{
				Err: "Malformed request",
			}, nil
		}

		if !state.Bootstrapped() {
			return reqres.SecretFetchResponse{
				Err: "safe not bootstrapped",
			}, nil
		}

		// TODO: sanitization end empty check.
		workloadId := r.WorkloadId
		workloadKey := r.WorkloadSecret

		// TODO: empty check
		if state.WorkloadKeyFromId(workloadId) != workloadKey {
			return reqres.SecretFetchResponse{
				Err: "I don’t know you, and it’s crazy…",
			}, nil
		}

		data, err := svc.SecretRead(ctx, workloadId)
		if err != nil {
			return reqres.SecretFetchResponse{
				Err: "Unknown problem fetching secret",
			}, nil
		}

		return reqres.SecretFetchResponse{
			Data: data,
		}, nil
	}
}
