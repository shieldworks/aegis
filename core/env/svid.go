/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                    <aegis.ist>
 *     .\_/.
 */

package env

import "os"

// SentinelSvidPrefix returns the prefix for the Safe
// SVID (Short-lived Verifiable Identity Document) used in the Aegis system.
// The prefix is obtained from the environment variable
// AEGIS_SENTINEL_SVID_PREFIX. If the variable is not set, the default prefix is
// used.
func SentinelSvidPrefix() string {
	p := os.Getenv("AEGIS_SENTINEL_SVID_PREFIX")
	if p == "" {
		p = "spiffe://aegis.ist/workload/aegis-sentinel/ns/aegis-system/sa/aegis-sentinel/n/"
	}
	return p
}

// SafeSvidPrefix returns the prefix for the Safe
// SVID (Short-lived Verifiable Identity Document) used in the Aegis system.
// The prefix is obtained from the environment variable
// AEGIS_SAFE_SVID_PREFIX. If the variable is not set, the default prefix is
// used.
func SafeSvidPrefix() string {
	p := os.Getenv("AEGIS_SAFE_SVID_PREFIX")
	if p == "" {
		p = "spiffe://aegis.ist/workload/aegis-safe/ns/aegis-system/sa/aegis-safe/n/"
	}
	return p
}

// NotarySvidPrefix returns the prefix for the Aegis Notary SVID
// (SPIFFE Verifiable Identity Document) used in the Aegis system.
// The prefix is obtained from the environment variable AEGIS_NOTARY_SVID_PREFIX.
// If the variable is not set, the default prefix is used.
//
// THIS IS NOT USED AT THE MOMENT.
// IT IS EXPERIMENTAL.
func NotarySvidPrefix() string {
	p := os.Getenv("AEGIS_NOTARY_SVID_PREFIX")
	if p == "" {
		p = "spiffe://aegis.ist/workload/aegis-notary/ns/aegis-system/sa/aegis-notary/n/"
	}
	return p
}

// WorkloadSvidPrefix returns the prefix for the Workload SVID
// (SPIFFE Verifiable Identity Document) used in the Aegis system.
// The prefix is obtained from the environment variable AEGIS_WORKLOAD_SVID_PREFIX.
// If the variable is not set, the default prefix is used.
func WorkloadSvidPrefix() string {
	p := os.Getenv("AEGIS_WORKLOAD_SVID_PREFIX")
	if p == "" {
		p = "spiffe://aegis.ist/workload/"
	}
	return p
}
