/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                  <aegis.z2h.dev>
 *     .\_/.
 */

package v1

import (
	"aegis-sidecar/internal/state"
	"context"
)

func (a apiV1Service) UpdateState(
	ctx context.Context, nextId, workloadId, workloadSecret, safeApiRoot string,
) error {
	state.Update(nextId, workloadId, workloadSecret, safeApiRoot)
	return nil
}
