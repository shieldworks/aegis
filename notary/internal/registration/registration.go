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
	"fmt"
	apiCoreV1 "k8s.io/api/core/v1"
	"os"
	"sync"
)

var mux sync.Mutex

type Payload struct {
	WorkloadId        string
	WorkloadSecret    string
	WorkloadNamespace string
	WorkloadName      string
	SafeApiRoot       string
	SafeBootstrapUrl  string
	WorkloadHookUrl   string
}

func Process(registrations chan (Payload)) {
	payload := <-registrations

	// Poor man’s transactional integrity:
	// Allow only one registration to be completed at a time.
	mux.Lock()
	defer mux.Unlock()

	notaryId := os.Getenv("AEGIS_NOTARY_ID")
	notaryToken := token.New()
	adminToken := token.New()
	safeBootstrapUrl := payload.SafeBootstrapUrl
	safeApiRoot := payload.SafeApiRoot
	workloadId := payload.WorkloadId
	workloadSecret := payload.WorkloadSecret

	workloadHookUrl := payload.WorkloadHookUrl

	err1 := bootstrap.Safe(safeBootstrapUrl, notaryId, notaryToken, adminToken)
	err2 := bootstrap.Sidecar(workloadHookUrl, notaryId, notaryToken, workloadId, workloadSecret, safeApiRoot)

	// TODO: if this thing fails more than N times, then discard it.
	// Use an exponential decay.
	if err1 == nil || err2 == nil {
		registrations <- payload
	}
}
func TryRegisterWorkload(registrations chan<- (Payload), pod *apiCoreV1.Pod) {
	if pod == nil {
		return
	}

	workloadKey, ok := pod.Annotations["aegis-workload-key"]
	if !ok {
		return
	}

	if workloadKey == "" {
		return
	}

	workloadId := pod.ObjectMeta.Name
	workloadSecret := token.New()
	workloadNamespace := pod.ObjectMeta.Namespace
	workloadName := pod.ObjectMeta.Name

	// TODO: we might want to make these configurable too.
	safeApiRoot := "http://aegis-safe.aegis-system.svc.cluster.local:8017/"
	safeBootstrapUrl := fmt.Sprintf("%sv1/bootstrap", safeApiRoot)
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
		WorkloadHookUrl:   workloadHookUrl,
	}
}

func TryUnregisterWorkload(registrations chan<- (Payload), pod *apiCoreV1.Pod) {
	// TODO: implement me!
}
