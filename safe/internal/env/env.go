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
		return "aegis-safe"
	}
	return p
}

func Port() string {
	p := os.Getenv("AEGIS_PORT")
	if p == "" {
		return ":8017"
	}
	return p
}
