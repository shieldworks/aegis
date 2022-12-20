/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                  <aegis.z2h.dev>
 *     .\_/.
 */

package v1

type WorkloadRegisterRequest struct {
	NotaryToken    string `json:"token"`
	WorkloadId     string `json:"workloadId"`
	WorkloadSecret string `json:"workloadSecret"`
	Err            string `json:"err,omitempty"`
}

type WorkloadRegisterResponse struct {
	Err string `json:"err,omitempty"`
}
