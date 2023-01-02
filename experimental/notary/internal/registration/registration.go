/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                  <aegis.z2h.dev>
 *     .\_/.
 */

package registration

import (
	"aegis-notary/internal/bootstrap"
	"aegis-notary/internal/crypto/token"
	"aegis-notary/internal/env"
	"fmt"
	apiCoreV1 "k8s.io/api/core/v1"
	"os"
)

// var mux sync.Mutex

type Payload struct {
	WorkloadId        string
	WorkloadSecret    string
	WorkloadNamespace string
	WorkloadName      string
	SafeApiRoot       string
	SafeBootstrapUrl  string
	SafeWorkloadUrl   string
	WorkloadHookUrl   string
}

func Process(registrations chan (Payload)) {
	payload := <-registrations

	if !env.InCluster() {
		// payload.SafeApiRoot = "http://localhost:8017/"
		payload.WorkloadHookUrl = "http://localhost:8039/v1/hook"
		payload.SafeBootstrapUrl = "http://localhost:8017/v1/bootstrap"
		payload.SafeWorkloadUrl = "http://localhost:8017/v1/workload"
	}

	notaryId := os.Getenv("AEGIS_NOTARY_ID")
	if notaryId == "" {
		notaryId = "AegisRocks"
	}

	// TODO: if there is a successful notary token that is registered, use that one; don't create a new one.
	notaryToken := token.New()

	// TODO: same for admin token.
	adminToken := token.New()

	// TODO: handle the case when `safe` pod is evicted:
	// 1. run the bootstrap flow again.
	// 2. safe may try to restore secrets
	// 3. safe can store its state on a persistent volume encrypted by a secret (see below discussion for how to secure that secret)

	// TODO: handle the case when `notary` is evicted
	// 1. notary will have to run the bootstrap flow again.
	// 2. since it is stateless, it will contact `safe` with its initial notary id, which will be rejected.
	// 3. notary can either encrypt and save notaryToken and adminToken to be able to restore upon crash, or just  bail out and lock itself.
	/*
				Access to that secret should be well-guarded.
				as per k8s documentation:
			In order to safely use Secrets, take at least the following steps:

			    Enable Encryption at Rest for Secrets.
			    Enable or configure RBAC rules with least-privilege access to Secrets.
			    Restrict Secret access to specific containers.
			    Consider using external Secret store providers.
		ref: https://kubernetes.io/docs/concepts/configuration/secret/
	*/

	safeBootstrapUrl := payload.SafeBootstrapUrl
	safeWorkloadUrl := payload.SafeWorkloadUrl
	safeApiRoot := payload.SafeApiRoot
	workloadId := payload.WorkloadId
	workloadSecret := payload.WorkloadSecret

	workloadHookUrl := payload.WorkloadHookUrl

	err1 := bootstrap.Safe(safeBootstrapUrl, notaryId, notaryToken, adminToken)

	/// TODO: if bootstrap fails, there is no need to do the next two.

	if err1 != nil {
		registrations <- payload
		return
	}

	err2 := bootstrap.Workload(workloadHookUrl, notaryId, notaryToken, workloadId, workloadSecret, safeApiRoot)
	err3 := bootstrap.SafeWorkload(safeWorkloadUrl, notaryToken, workloadId, workloadSecret)

	/// err3 := bootstrap.SafeWorkload()
	// http://aegis-safe.aegis-system.svc.cluster.local:8017/v1/workload

	// TODO: if this thing fails more than N times, then discard it.
	// Use an exponential decay.
	if err2 != nil || err3 != nil {
		registrations <- payload
	}
}

func TryRegisterWorkload(registrations chan<- (Payload), pod *apiCoreV1.Pod) {
	if pod == nil {
		return
	}

	wi, ok := pod.Annotations["aegis-workload-id"]
	if !ok {
		return
	}

	if wi == "" {
		return
	}

	workloadId := wi // pod.ObjectMeta.Name
	workloadSecret := token.New()
	workloadNamespace := pod.ObjectMeta.Namespace
	workloadName := pod.ObjectMeta.Name

	// TODO: we might want to make these configurable too.
	safeApiRoot := "http://aegis-safe.aegis-system.svc.cluster.local:8017/"
	safeBootstrapUrl := fmt.Sprintf("%sv1/bootstrap", safeApiRoot)
	safeWorkloadUrl := fmt.Sprintf("%sv1/workspace", safeApiRoot)
	workloadHookUrl := fmt.Sprintf(
		"http://%s.%s.svc.cluster.local:8039/v1/hook",
		workloadName, workloadNamespace,
	)

	registrations <- Payload{
		WorkloadId:        workloadId,
		WorkloadSecret:    workloadSecret,
		WorkloadNamespace: workloadNamespace,
		WorkloadName:      workloadName,
		SafeApiRoot:       safeApiRoot,
		SafeBootstrapUrl:  safeBootstrapUrl,
		SafeWorkloadUrl:   safeWorkloadUrl,
		WorkloadHookUrl:   workloadHookUrl,
	}
}

func TryUnregisterWorkload(registrations chan<- (Payload), pod *apiCoreV1.Pod) {
	// TODO: implement me!
}
