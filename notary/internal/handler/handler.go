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

func ListenPodEvents(podInformer informersCoreV1.PodInformer, registrations chan registration.Payload) error {
	_, err := podInformer.Informer().AddEventHandler(
		cache.ResourceEventHandlerFuncs{
			AddFunc: func(pod interface{}) {
				p := pod.(*apiCoreV1.Pod)
				if !meta.AegisAnnotatedWorkload(p) {
					return
				}
				registration.TryRegisterWorkload(registrations, p)
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
