/*
 * .-'_.---._'-.
 * ||####|(__)||   Protect your secrets, protect your business.
 *   \\()|##//       Secure your sensitive data with Aegis.
 *    \\ |#//                  <aegis.z2h.dev>
 *     .\_/.
 */

package main

import (
	"aegis-notary/internal/handler"
	"aegis-notary/internal/registration"
	"flag"
	"fmt"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"os"
	"time"
)

func main() {
	registrations := make(chan registration.Payload)
	go registration.Process(registrations)

	// TODO: this probably need to change for in-cluster deployment as an operator.
	// Read in and parse Kubernetes config
	kubeconfig := flag.String(
		"kubeconfig", os.Getenv("HOME")+"/.kube/config", "kubeconfig file",
	)

	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// Generate a new 'Clientset' which is used to interact with the Kubernetes API
	clientset, createClientErr := kubernetes.NewForConfig(config)
	if createClientErr != nil {
		panic(createClientErr)
	}

	// TODO: increase this interval.
	// TODO: make it configurable.
	// Generate an Informer factory so that we can listen in to built-in resources
	informerFactory := informers.NewSharedInformerFactory(clientset, time.Second*1)

	podInformer := informerFactory.Core().V1().Pods()

	err = handler.ListenPodEvents(podInformer, registrations)

	if err != nil {
		fmt.Println("poop")
		return
	}

	stop := make(chan struct{})
	defer close(stop)

	// Start our informers
	informerFactory.Start(wait.NeverStop)
	informerFactory.WaitForCacheSync(wait.NeverStop)

	fmt.Println("everything is awesome!")
	select {} // block forever.
}
