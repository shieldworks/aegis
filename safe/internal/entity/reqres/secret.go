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

type SecretFetchRequest struct {
	WorkloadToken string `json:"token"`
	WorkloadId    string `json:"id"`
	WorkloadKey   string `json:"key"`
	Err           string `json:"err,omitempty"`
}

type SecretFetchResponse struct {
	Data string `json:"data"`
	Err  string `json:"err,omitempty"`
}
