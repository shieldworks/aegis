package main

import (
	"context"
	"flag"
	"fmt"

	"os"
	"time"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"

	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {

	// Read in and parse Kubernetes config
	kubeconfig := flag.String("kubeconfig", os.Getenv("HOME")+"/.kube/config", "kubeconfig file")

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

	// Generate an Informer factory so that we can listen in to built-in resources
	informerFactory := informers.NewSharedInformerFactory(clientset, time.Second*20)

	configMapInformer := informerFactory.Core().V1().ConfigMaps()
	podInformer := informerFactory.Core().V1().Pods()

	// Add event handlers!
	// This constitutes the "Read" phase of the control loop
	configMapInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{

		// Every time a ConfigMap with the label 'app=test' is created, delete the 'test' pod.
		AddFunc: func(obj interface{}) {
			configMap := obj.(*corev1.ConfigMap)
			// Check specific cluster state. Here we are in the "Business Logic" or "Diffing" phase of the control loop
			if value, ok := configMap.Labels["app"]; ok {
				if value == "test" {
					fmt.Println("Found config map with label app=test")

					// Delete the 'test' pod - we are altering the state of the cluster. Here we are in the "Update" phase of the control loop
					deleteErr := clientset.CoreV1().Pods("default").Delete(context.TODO(), "test", metav1.DeleteOptions{})

					if deleteErr != nil {
						panic(err)
					} else {
						fmt.Printf("Deleted pod 'test'\n")

					}
				}
			}
		},
	})

	podInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{

		// Any time the 'test' pod is deleted - recreate it!
		DeleteFunc: func(obj interface{}) {
			deletedPod := obj.(*corev1.Pod)
			if deletedPod.Name == "test" {
				// Note: This is a way to generate the config dynamically in code - you can also read this config from a yaml file.
				pod := corev1.Pod{
					ObjectMeta: metav1.ObjectMeta{
						Name: "test",
						Labels: map[string]string{
							"app": "test",
						},
					},
					Spec: corev1.PodSpec{
						Containers: []corev1.Container{
							{
								Name:  "nginx",
								Image: "nginx",
							},
						},
					},
				}

				// Create the pod
				createdPod, createErr := clientset.CoreV1().Pods("default").Create(context.TODO(), &pod, metav1.CreateOptions{})

				if createErr != nil {
					fmt.Printf("Error recreating pod: %s\n", createErr)
					panic(createErr)
				} else {
					fmt.Printf("Rereated pod: %s\n", createdPod.Name)
				}
			}

		},
	})

	stop := make(chan struct{})
	defer close(stop)

	// Start our informers
	informerFactory.Start(wait.NeverStop)
	informerFactory.WaitForCacheSync(wait.NeverStop)

	// Wait on the 'stop' channel. This let's this process continue on running.
	// You can stop it with 'CTRL-c' when running directly from a terminal
	select {}
}
