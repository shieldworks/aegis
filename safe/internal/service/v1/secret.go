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
	ctx context.Context, key, value string,
) error {
	state.UpsertSecret(key, value)
	return nil
}

func (a apiV1Service) SecretRead(
	ctx context.Context, key string,
) (string, error) {
	val := state.ReadSecret(key)
	return val, nil
}
