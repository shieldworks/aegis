/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                  <aegis.z2h.dev>
 *     .\_/.
 */

package bootstrap

import (
	reqres "aegis-notary/internal/entity/reqres/v1"
	"aegis-notary/internal/upstream"
	"context"
	"sync"
)

func SafeWorkload(safeWorkloadUrl, notaryToken, workloadId, workloadSecret string) error {
	//fmt.Println("bootstrap.SafeWorkload",
	//	"safeWorkloadUrl", safeWorkloadUrl,
	//	"notaryToken", notaryToken,
	//	"workloadId", workloadId,
	//	"workloadSecret", workloadSecret)

	_, err := upstream.NewSafeWorkloadRegisterEndpoint(safeWorkloadUrl)(
		context.Background(),
		reqres.WorkloadRegisterRequest{
			NotaryToken:    notaryToken,
			WorkloadId:     workloadId,
			WorkloadSecret: workloadSecret,
		},
	)

	if err != nil {
		// TODO: handle me
		panic("handle me")
	}

	// TODO: check status code too.

	return nil
}

func Workload(workloadHookUrl, notaryId, newNotaryId, workloadId, workloadSecret, safeApiRoot string) error {
	//fmt.Println("bootstrap.Workload",
	//	"url", workloadHookUrl,
	//	"notaryId", notaryId,
	//	"newNotaryId", newNotaryId,
	//	"workloadId", workloadId,
	//	"workloadSecret", workloadSecret,
	//	"safeApiRoot", safeApiRoot,
	//)

	_, err := upstream.NewSidecarBootstrapEndpoint(workloadHookUrl)(
		context.Background(),
		reqres.HookRequest{
			NotaryId:       notaryId,
			NewNotaryId:    newNotaryId,
			WorkloadId:     workloadId,
			WorkloadSecret: workloadSecret,
			SafeApiRoot:    safeApiRoot,
		})

	if err != nil {
		// TODO: handle me
		panic("handle me")
	}

	// TODO: check status code too.

	return nil
}

// TODO: this is not ideal. notary should periodically check safe’s bootstrap status
// and bootstrap it if it has not been bootstrapped.
// TODO: notary should periodically rotate safe’s admin token too.
// TODO: safe needs a CLI to get the current admin token.
var mux sync.Mutex
var safeBootstrapped = false

func Safe(safeBootStrapUrl, notaryId, notaryToken, adminToken string) error {
	//fmt.Println("bootstrap.Safe",
	//	"url", safeBootStrapUrl,
	//	"notaryId", notaryId,
	//	"notaryToken", notaryToken,
	//	"adminToken", adminToken)

	mux.Lock()
	if safeBootstrapped {
		mux.Unlock()
		return nil
	}
	mux.Unlock()

	_, err := upstream.NewSafeBootstrapEndpoint(safeBootStrapUrl)(
		context.Background(),
		reqres.BootstrapRequest{
			NotaryId:    notaryId,
			NotaryToken: notaryToken,
			AdminToken:  adminToken,
		})

	if err != nil {
		// TODO: handle me
		panic("handle me")
	}

	mux.Lock()
	defer mux.Unlock()
	// TODO: check status code too.
	safeBootstrapped = true
	return nil
}
