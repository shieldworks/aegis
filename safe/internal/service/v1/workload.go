/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                  <aegis.z2h.dev>
 *     .\_/.
 */

package v1

import (
	"aegis-safe/internal/state"
	"context"
)

func (a apiV1Service) WorkloadRegister(ctx context.Context, id, secret string) error {
	state.RegisterWorkload(id, secret)
	return nil
}
