package v1

/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                  <aegis.z2h.dev>
 *     .\_/.
 */

import (
	"aegis-safe/internal/state"
	"context"
)

func (a apiV1Service) Bootstrap(
	ctx context.Context, adminToken, workloadToken string,
) error {
	if state.Bootstrapped() {
		return nil
	}

	state.Bootstrap(adminToken, workloadToken)

	return nil
}
