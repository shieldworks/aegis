/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                  <aegis.z2h.dev>
 *     .\_/.
 */

package v1

import "context"

func (a apiV1Service) SecretUpsert(
	ctx context.Context, key, value string,
) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (a apiV1Service) SecretRead(
	ctx context.Context, key string,
) (string, error) {
	//TODO implement me
	panic("implement me")
}
