package state

/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                    <aegis.ist>
 *     .\_/.
 */

import "strings"

func ageKeyTriplet() (string, string, string) {
	if ageKey == "" {
		return "", "", ""
	}

	parts := strings.Split(ageKey, "\n")

	if len(parts) != 3 {
		return "", "", ""
	}

	return parts[0], parts[1], parts[2]
}
