/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                    <aegis.ist>
 *     .\_/.
 */

package env

import "os"

// SafeEndpointUrl returns the URL for the Aegis Safe endpoint used in the Aegis system.
// The URL is obtained from the environment variable AEGIS_SAFE_ENDPOINT_URL.
// If the variable is not set, the default URL is used.
func SafeEndpointUrl() string {
	u := os.Getenv("AEGIS_SAFE_ENDPOINT_URL")
	if u == "" {
		u = "https://aegis-safe.aegis-system.svc.cluster.local:8443/"
	}
	return u
}

// NotaryEndpointUrl returns the URL for the Aegis Notary endpoint used in the
// Aegis system. The URL is obtained from the environment variable
// AEGIS_NOTARY_ENDPOINT_URL. If the variable is not set, the default URL is used.
//
// THIS IS NOT BEING USED. IT IS EXPERIMENTAL.
func NotaryEndpointUrl() string {
	u := os.Getenv("AEGIS_NOTARY_ENDPOINT_URL")
	if u == "" {
		u = "https://aegis-notary.aegis-system.svc.cluster.local:8443/"
	}
	return u
}
