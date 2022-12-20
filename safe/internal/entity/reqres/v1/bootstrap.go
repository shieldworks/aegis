/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                  <aegis.z2h.dev>
 *     .\_/.
 */

package v1

type BootstrapRequest struct {
	NotaryId    string `json:"id"`
	NotaryToken string `json:"notaryToken"`
	AdminToken  string `json:"adminToken"`
	Err         string `json:"err,omitempty"`
}

type BootstrapResponse struct {
	Err string `json:"err,omitempty"`
}
