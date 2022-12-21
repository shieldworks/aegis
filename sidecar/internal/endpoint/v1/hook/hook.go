/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                  <aegis.z2h.dev>
 *     .\_/.
 */

package hook

import (
	reqres "aegis-sidecar/internal/entity/reqres/v1"
	service "aegis-sidecar/internal/service/v1"
	"aegis-sidecar/internal/state"
	"context"
	"github.com/go-kit/kit/endpoint"
	"os"
)

func MakeHookEndpoint(svc service.ApiV1Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		r, ok := request.(reqres.HookRequest)
		if !ok {
			return reqres.HookResponse{
				Err: "malformed payload",
			}, nil
		}

		// Already bootstrapped; discard request.
		// Will need to change it if/when we implement token rotation.
		if state.Bootstrapped() {
			return reqres.HookResponse{}, nil
		}

		// TODO: sanitize these.
		notaryId := r.NotaryId
		nextId := r.NewNotaryId
		workloadId := r.WorkloadId
		workloadSecret := r.WorkloadSecret
		safeApiRoot := r.SafeApiRoot

		// TODO: empty check
		envNotaryId := os.Getenv("AEGIS_NOTARY_ID")

		if envNotaryId != notaryId {
			return reqres.HookResponse{
				Err: "I don’t know you, and it’s crazy…",
			}, nil
		}

		err := svc.UpdateState(ctx, nextId, workloadId, workloadSecret, safeApiRoot)
		if err != nil {
			return reqres.HookResponse{
				Err: "Unknown problem while processing hook.",
			}, nil
		}

		return reqres.HookResponse{}, nil
	}
}
