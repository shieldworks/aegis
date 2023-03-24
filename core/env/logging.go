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
// 1: logs are off, 6: highest verbosity.
// Off = 1, Error = 2, Warn = 3, Info = 4, Debug = 5, Trace = 6
func LogLevel() int {
	p := os.Getenv("AEGIS_LOG_LEVEL")
	if p == "" {
		return 3
	}
	l, _ := strconv.Atoi(p)
	if l == 0 {
		return 3
	}
	if l < 0 || l > 6 {
		return 3
	}
	return l
}
