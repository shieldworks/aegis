/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                  <aegis.z2h.dev>
 *     .\_/.
 */

package reqres

type BootstrapRequest struct {
	NotaryId      string `json:"id"`
	WorkloadToken string `json:"workloadToken"`
	AdminToken    string `json:"adminToken"`
	Err           string `json:"err,omitempty"`
}

type BootstrapResponse struct {
	Err string `json:"err,omitempty"`
}
