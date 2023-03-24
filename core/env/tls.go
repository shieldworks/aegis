/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                    <aegis.ist>
 *     .\_/.
 */

package env

import "os"

// TlsPort returns the secure port for Aegis Safe to listen on.
// It checks the AEGIS_TLS_PORT environment variable. If the variable
// is not set, it defaults to ":8443".
func TlsPort() string {
	p := os.Getenv("AEGIS_TLS_PORT")
	if p == "" {
		p = ":8443"
	}
	return p
}
