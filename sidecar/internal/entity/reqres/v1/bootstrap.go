/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                  <aegis.z2h.dev>
 *     .\_/.
 */

package v1

type BootstrapRequest struct {
	NotaryToken    string `json:"token"`
	NotarySecret   string `json:"notarySecret"`
	WorkloadId     string `json:"workloadId"`
	WorkloadSecret string `json:"workloadSecret"`
	SafeApiRoot    string `json:"safeApiRoot"`
	Err            string `json:"err,omitempty"`
}

type BootstrapResponse struct {
	Err string `json:"err,omitempty"`
}
