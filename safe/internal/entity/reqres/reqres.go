/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                  <aegis.z2h.dev>
 *     .\_/.
 */

package reqres

type SecretUpsertRequest struct {
	AdminToken string `json:"token"`
	Key        string `json:"key"`
	Value      string `json:"value"`
	Err        string `json:"err,omitempty"`
}

type SecretUpsertResponse struct {
	Err string `json:"err,omitempty"`
}

type BootstrapRequest struct {
	NotaryId      string `json:"id"`
	WorkloadToken string `json:"workloadToken"`
	AdminToken    string `json:"adminToken"`
	Err           string `json:"err,omitempty"`
}

type BootstrapResponse struct {
	Err string `json:"err,omitempty"`
}

type SecretReadRequest struct {
	Err string `json:"err,omitempty"`
}

type SecretReadResponse struct {
	Err string `json:"err,omitempty"`
}

type GenericRequest struct {
	Err string `json:"err,omitempty"`
}

type GenericResponse struct {
	Err string `json:"err,omitempty"`
}
