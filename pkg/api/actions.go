package api

import (
	"context"
	"fmt"
	"log"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func getPods(cl *kubernetes.Clientset, ns string) *v1.PodList {

	// Retrieve the Corev1 Client via clientset and list all Nodes in the cluster
	pods, err := cl.CoreV1().Pods(ns).List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		log.Fatal(err)
	}

	return pods

}

func getPvcs(cl *kubernetes.Clientset, ns string) *v1.PersistentVolumeClaimList {

	// Possible LabelSelector and FieldSelector
	// listOptions := metav1.ListOptions{LabelSelector: label, FieldSelector: field}

	pvcs, err := cl.CoreV1().PersistentVolumeClaims(ns).List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println()

	return pvcs
}

func getServices(cl *kubernetes.Clientset, ns string) *v1.ServiceList {
	svcs, err := cl.CoreV1().Services(ns).List(context.TODO(), metav1.ListOptions{})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println()
	return svcs
}
