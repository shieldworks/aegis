/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                  <aegis.z2h.dev>
 *     .\_/.
 */

package env

import "os"

func AppName() string {
	p := os.Getenv("AEGIS_APP_NAME")
	if p == "" {
		return "aegis-sidecar"
	}
	return p
}

func Port() string {
	p := os.Getenv("AEGIS_PORT")
	if p == "" {
		return ":8039"
	}
	return p
}

func SafeApiRoot() string {
	p := os.Getenv("AEGIS_SAFE_API_ROOT")
	if p == "" {
		return "http://aegis-safe.aegis-system.svc.cluster.local:8017/"
	}
	return p
}
