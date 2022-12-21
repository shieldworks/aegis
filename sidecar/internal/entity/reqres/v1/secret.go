/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                  <aegis.z2h.dev>
 *     .\_/.
 */

package v1

// TODO: these entities are repeated, so they may go to a common library maybe.

type SecretFetchRequest struct {
	WorkloadId     string `json:"workload"`
	WorkloadSecret string `json:"secret"`
	Err            string `json:"err,omitempty"`
}

type SecretFetchResponse struct {
	Data string `json:"data"`
	Err  string `json:"err,omitempty"`
}
