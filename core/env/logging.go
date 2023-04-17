/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                    <aegis.ist>
 *     .\_/.
 */

package env

import (
	"os"
	"strconv"
)

// LogLevel returns the value set by AEGIS_LOG_LEVEL environment
// variable, or a default level.
//
// AEGIS_LOG_LEVEL determines the verbosity of the logs.
// 0: logs are off, 7: highest verbosity (TRACE).
func LogLevel() int {
	p := os.Getenv("AEGIS_LOG_LEVEL")
	if p == "" {
		return int(3) // WARN
	}
	l, _ := strconv.Atoi(p)
	if l == 0 {
		return 3 // WARN
	}
	if l < 0 || l > 7 {
		return 3 // WARN
	}
	return l
}
