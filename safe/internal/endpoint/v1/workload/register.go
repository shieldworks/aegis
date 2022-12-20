/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                  <aegis.z2h.dev>
 *     .\_/.
 */

package workload

import (
	"aegis-safe/internal/entity/reqres"
	service "aegis-safe/internal/service/v1"
	"aegis-safe/internal/state"
	"context"
	"github.com/go-kit/kit/endpoint"
)

func MakeWorkloadRegisterEndpoint(svc service.ApiV1Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		r, ok := request.(reqres.WorkloadRegisterRequest)
		if !ok {
			return reqres.WorkloadRegisterResponse{
				Err: "malformed request",
			}, nil
		}

		if !state.Bootstrapped() {
			return reqres.WorkloadRegisterResponse{
				Err: "safe not bootstrapped",
			}, nil
		}

		// TODO: sanitization, empty check, etc.
		workloadToken := r.WorkloadToken
		workloadId := r.WorkloadId
		workloadSecret := r.WorkloadSecret

		if workloadToken != state.NotaryWorkloadToken() {
			return reqres.WorkloadRegisterResponse{
				Err: "I don’t know you, and it’s crazy…",
			}, nil
		}

		err := svc.WorkloadRegister(ctx, workloadId, workloadSecret)
		if err != nil {
			return reqres.WorkloadRegisterResponse{
				Err: "Unknown problem registering workload",
			}, nil
		}

		return reqres.WorkloadRegisterResponse{}, nil
	}
}
