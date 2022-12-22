/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                  <aegis.z2h.dev>
 *     .\_/.
 */

package v1

type HookRequest struct {
	NotaryId       string `json:"id"`
	NewNotaryId    string `json:"nextId"`
	WorkloadId     string `json:"workloadId"`
	WorkloadSecret string `json:"workloadSecret"`
	SafeApiRoot    string `json:"safeApiRoot"`
	Err            string `json:"err,omitempty"`
}

type HookResponse struct {
	Err string `json:"err,omitempty"`
}

type BootstrapRequest struct {
	NotaryId    string `json:"id"`
	NotaryToken string `json:"notaryToken"`
	AdminToken  string `json:"adminToken"`
	Err         string `json:"err,omitempty"`
}

type BootstrapResponse struct {
	Err string `json:"err,omitempty"`
}
