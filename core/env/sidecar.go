/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                    <aegis.ist>
 *     .\_/.
 */

package env

import "os"

// SidecarSecretsPath returns the path to the secrets file used by the sidecar.
// The path is determined by the AEGIS_SIDECAR_SECRETS_PATH environment variable,
// with a default value of "/opt/aegis/secrets.json" if the variable is not set.
func SidecarSecretsPath() string {
	p := os.Getenv("AEGIS_SIDECAR_SECRETS_PATH")
	if p == "" {
		p = "/opt/aegis/secrets.json"
	}
	return p
}
