/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                  <aegis.z2h.dev>
 *     .\_/.
 */

package reqres

type SecretUpsertRequest struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	Err   string `json:"err,omitempty"`
}

type SecretUpsertResponse struct {
	Token string `json:"token"`
	Err   string `json:"err,omitempty"`
}

type BootstrapRequest struct {
	NotaryId string `json:"notaryId"`
	Token    string `json:"token"`
	Err      string `json:"err,omitempty"`
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
