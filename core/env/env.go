/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                  <aegis.z2h.dev>
 *     .\_/.
 */

package env

import "os"

func SentinelSvidPrefix() string {
	p := os.Getenv("AEGIS_SENTINEL_SVID_PREFIX")
	if p == "" {
		p = "spiffe://aegis.z2h.dev/workload/aegis-sentinel/ns/aegis-system/sa/aegis-sentinel/n/"
	}
	return p
}

func SafeSvidPrefix() string {
	p := os.Getenv("AEGIS_SAFE_SVID_PREFIX")
	if p == "" {
		p = "spiffe://aegis.z2h.dev/workload/aegis-safe/ns/aegis-system/sa/aegis-safe/n/"
	}
	return p
}

func WorkloadSvidPrefix() string {
	p := os.Getenv("AEGIS_WORKLOAD_SVID_PREFIX")
	if p == "" {
		p = "spiffe://aegis.z2h.dev/workload/"
	}
	return p
}
