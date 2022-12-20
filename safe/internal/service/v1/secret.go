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

func (a apiV1Service) SecretUpsert(
	ctx context.Context, id, data string,
) error {
	state.UpsertSecret(id, data)
	return nil
}

func (a apiV1Service) SecretRead(
	ctx context.Context, id string,
) (string, error) {
	data := state.ReadSecret(id)
	return data, nil
}
