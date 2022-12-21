/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                  <aegis.z2h.dev>
 *     .\_/.
 */

package v1

import "context"

type ApiV1Service interface {
	UpdateState(ctx context.Context, nextId, workloadId, workloadSecret, safeApiRoot string) error
}

type apiV1Service struct{}

func NewApiV1Service() ApiV1Service {
	return apiV1Service{}
}
