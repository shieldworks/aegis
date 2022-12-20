/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                  <aegis.z2h.dev>
 *     .\_/.
 */

package reqres

type GenericRequest struct {
	Err string `json:"err,omitempty"`
}

type GenericResponse struct {
	Err string `json:"err,omitempty"`
}
