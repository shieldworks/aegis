/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                  <aegis.z2h.dev>
 *     .\_/.
 */

package handler

import (
	"aegis-notary/internal/meta"
	"aegis-notary/internal/registration"
	apiCoreV1 "k8s.io/api/core/v1"
	informersCoreV1 "k8s.io/client-go/informers/core/v1"
	"k8s.io/client-go/tools/cache"
)

// TODO: OnUpdate is called with all existing objects on the specific resync interval
// This can be especially useful if a Pod hasn’t been registered for whatever reason.
// we can keep an in-memory dictionary of registered workloads, and if the workload
// in question is not in the list, we can enqueue it to be registered next.

func ListenPodEvents(podInformer informersCoreV1.PodInformer, registrations chan registration.Payload) error {
	_, err := podInformer.Informer().AddEventHandler(
		cache.ResourceEventHandlerFuncs{
			AddFunc: func(pod interface{}) {
				p := pod.(*apiCoreV1.Pod)
				if !meta.AegisAnnotatedWorkload(p) {
					return
				}
				// fmt.Println("AAAA", p.Name)
				registration.TryRegisterWorkload(registrations, p)
			},
			UpdateFunc: func(oldPod interface{}, newPod interface{}) {
				// op := oldPod.(*apiCoreV1.Pod)
				np := newPod.(*apiCoreV1.Pod)

				if meta.AegisAnnotatedWorkload(np) {
					/// fmt.Println("BBBBBB", np.Name)
					// registration.TryRegisterWorkload(registrations, np)
				}
			},
			DeleteFunc: func(pod interface{}) {
				p := pod.(*apiCoreV1.Pod)
				if !meta.AegisAnnotatedWorkload(p) {
					return
				}
				registration.TryUnregisterWorkload(registrations, p)
			},
		},
	)
	return err
}
