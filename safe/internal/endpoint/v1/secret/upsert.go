/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                  <aegis.z2h.dev>
 *     .\_/.
 */

package secret

import (
	"aegis-safe/internal/entity/reqres"
	service "aegis-safe/internal/service/v1"
	"aegis-safe/internal/state"
	"context"
	"github.com/go-kit/kit/endpoint"
)

func MakeSecretUpsertEndpoint(svc service.ApiV1Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		r, ok := request.(reqres.SecretUpsertRequest)
		if !ok {
			return reqres.SecretUpsertResponse{
				Err: "malformed request",
			}, nil
		}

		if !state.Bootstrapped() {
			return reqres.SecretUpsertResponse{
				Err: "safe not bootstrapped",
			}, nil
		}

		// TODO: sanitize these
		// also empty check.
		adminToken := r.AdminToken
		key := r.Key
		value := r.Value

		if adminToken != state.NotaryAdminToken() {
			return reqres.SecretUpsertResponse{
				Err: "I don’t know you, and it’s crazy…",
			}, nil
		}

		err := svc.SecretUpsert(ctx, key, value)
		if err != nil {
			return reqres.SecretUpsertResponse{
				Err: "Unknown problem while adding secret",
			}, nil
		}

		return reqres.SecretUpsertResponse{}, nil
	}
}
