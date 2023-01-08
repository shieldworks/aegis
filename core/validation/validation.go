/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                  <aegis.z2h.dev>
 *     .\_/.
 */

package validation

import (
	"github.com/zerotohero-dev/aegis-core/env"
	"strings"
)

func IsSentinel(svid string) bool {
	return strings.HasPrefix(svid, env.SentinelSvidPrefix())
}

func IsSafe(svid string) bool {
	return strings.HasPrefix(svid, env.SafeSvidPrefix())
}

func IsWorkload(svid string) bool {
	return strings.HasPrefix(svid, env.WorkloadSvidPrefix())
}
