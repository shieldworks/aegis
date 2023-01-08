/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                  <aegis.z2h.dev>
 *     .\_/.
 */

package env

import (
	"os"
	"strconv"
	"time"
)

func SpiffeSocketUrl() string {
	p := os.Getenv("SPIFFE_ENDPOINT_SOCKET")
	if p == "" {
		p = "unix:///spire-agent-socket/agent.sock"
	}
	return p
}

func SafeEndpointUrl() string {
	u := os.Getenv("AEGIS_SAFE_ENDPOINT_URL")
	if u == "" {
		u = "https://aegis-safe.aegis-system.svc.cluster.local:8443/"
	}
	return u
}

func SidecarSecretsPath() string {
	p := os.Getenv("AEGIS_SIDECAR_SECRETS_PATH")
	if p == "" {
		p = "/opt/aegis/secrets.json"
	}
	return p
}

func SentryPollInterval() time.Duration {
	p := os.Getenv("AEGIS_SIDECAR_POLL_INTERVAL")
	if p == "" {
		p = "20"
	}
	i, err := strconv.ParseInt(p, 10, 32)
	if err != nil {
		return 20 * time.Second
	}
	return time.Duration(i) * time.Second
}
