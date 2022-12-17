package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/informers"
	k8s "k8s.io/client-go/kubernetes"
	listersv1 "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
)

/**
* The code that follows was inspired by the custom scheduler found here:
* https://github.com/martonsereg/random-scheduler/blob/master/cmd/scheduler/main.go
*
* This simple scheduler will quickly assign a Pod to the Node which has the least amount of Pods
* currently residing upon it.
**/

type CustomScheduler struct {
	clientset       *k8s.Clientset
	podQueueChannel chan *v1.Pod
	nodeLister      listersv1.NodeLister
}

var schedulerName = "custom-scheduler"

// Generate a Kubernetes Clientset
func generateClientset() *k8s.Clientset {

	// Read in and parse Kubernetes config
	kubeconfig := flag.String("kubeconfig", os.Getenv("HOME")+"/.kube/config", "kubeconfig file")

	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	clientset, err := k8s.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	return clientset
}

// Finds the best fitting node for the current pod in the queue and binds the pod to it.
func (s *CustomScheduler) schedule() {
	pod := <-s.podQueueChannel
	fmt.Println("Pod to schedule:", pod.Namespace, "/", pod.Name)

	nodes, err := s.nodeLister.List(labels.Everything())
	if err != nil || len(nodes) < 1 {
		fmt.Println("Couldn't find available nodes.", err)
		return
	}

	// 2. Find the best fitting Node for the available Pod
	bestFitCount := -1
	bestFitNode := nodes[0]
	firstNodePods, _ := s.clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{
		FieldSelector: "spec.nodeName=" + bestFitNode.Name,
	})

	if firstNodePods.Size() < bestFitCount {
		bestFitCount = firstNodePods.Size()
	}

	// Iterate through every node - find the node with the least amount of pods on it
	for _, node := range nodes[1:] {
		nodeName := node.Name
		pods, err := s.clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{
			FieldSelector: "spec.nodeName=" + nodeName,
		})
		if err != nil {
			continue
		}

		if pods.Size() < bestFitCount {
			bestFitCount = pods.Size()
			bestFitNode = node.DeepCopy()
		}
	}

	// 3. Bind the Pod to the Node that best fits.
	s.clientset.CoreV1().Pods(pod.Namespace).Bind(context.TODO(), &v1.Binding{
		ObjectMeta: metav1.ObjectMeta{
			Name:      pod.Name,
			Namespace: pod.Namespace,
		},
		Target: v1.ObjectReference{
			APIVersion: "v1",
			Kind:       "Node",
			Name:       bestFitNode.Name,
		},
	}, metav1.CreateOptions{})

	eventMsg := fmt.Sprintf("Bound pod %s.%s on %s\n", pod.Namespace, pod.Name, bestFitNode.Name)
	fmt.Println(eventMsg)

	// Shoot out an event - can consume with 'kubectl get events'
	now := time.Now().UTC()
	_, eventErr := s.clientset.CoreV1().Events(pod.Namespace).Create(context.TODO(), &v1.Event{
		Count:          1,
		Message:        eventMsg,
		Reason:         "Scheduled",
		LastTimestamp:  metav1.NewTime(now),
		FirstTimestamp: metav1.NewTime(now),
		Type:           "Normal",
		Source: v1.EventSource{
			Component: schedulerName,
		},
		InvolvedObject: v1.ObjectReference{
			Kind:      "Pod",
			Name:      pod.Name,
			Namespace: pod.Namespace,
			UID:       pod.UID,
		},
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: pod.Name + "-",
		},
	}, metav1.CreateOptions{})

	if eventErr != nil {
		fmt.Println("Error creating event!", eventErr)
	}

}

func main() {

	// Start a channel that will house the current, unassigned Pod
	podQueueChannel := make(chan *v1.Pod, 1000)
	defer close(podQueueChannel)

	// Start a channel that will eventually close the program
	stopChannel := make(chan struct{})
	defer close(stopChannel)

	clientset := generateClientset()

	// 1. Watch Pods and Nodes
	informerFactory := informers.NewSharedInformerFactory(clientset, 200)

	nodeInformer := informerFactory.Core().V1().Nodes()
	nodeInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			node, ok := obj.(*v1.Node)
			if !ok {
				return
			}
			log.Printf("New node being tracked by custom-scheduler node informer: %s", node.GetName())
		},
	})

	podInformer := informerFactory.Core().V1().Pods()
	podInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			pod, ok := obj.(*v1.Pod)
			if !ok {
				return
			}

			// Remember to set the scheduler name in the Pod's manifest when running multiple schedulers
			// If the newly added pod needs to be scheduled using this scheduler then proceed, otherwise -
			// we'll let kube-scheduler do its thing.
			if pod.Spec.NodeName == "" && pod.Spec.SchedulerName == schedulerName {
				podQueueChannel <- pod
			}
		},
	})

	informerFactory.Start(stopChannel)

	nodeLister := nodeInformer.Lister()

	scheduler := CustomScheduler{
		nodeLister:      nodeLister,
		podQueueChannel: podQueueChannel,
		clientset:       clientset,
	}

	wait.Until(scheduler.schedule, 0, stopChannel)
}
