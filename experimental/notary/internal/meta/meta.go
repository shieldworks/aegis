/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                  <aegis.z2h.dev>
 *     .\_/.
 */

package meta

import apiCoreV1 "k8s.io/api/core/v1"

func AegisAnnotatedWorkload(pod *apiCoreV1.Pod) bool {
	if pod == nil {
		return false
	}

	workloadId, ok := pod.Annotations["aegis-workload-id"]
	if !ok {
		return false
	}

	return workloadId != ""
}
