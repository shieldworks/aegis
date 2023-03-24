/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                    <aegis.ist>
 *     .\_/.
 */

package env

import "os"

// ProbeLivenessPort returns the port for liveness probe.
// It first checks the environment variable AEGIS_PROBE_LIVENESS_PORT.
// If the variable is not set, it returns the default value ":8081".
func ProbeLivenessPort() string {
	u := os.Getenv("AEGIS_PROBE_LIVENESS_PORT")
	if u == "" {
		u = ":8081"
	}
	return u
}

// ProbeReadinessPort returns the port for readiness probe.
// It first checks the environment variable AEGIS_PROBE_READINESS_PORT.
// If the variable is not set, it returns the default value ":8082".
func ProbeReadinessPort() string {
	u := os.Getenv("AEGIS_PROBE_READINESS_PORT")
	if u == "" {
		u = ":8082"
	}
	return u
}
