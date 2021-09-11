package main

import (
	"context"
	"log"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func getPods(cl *kubernetes.Clientset, ns string) *v1.PodList {

	// Retrieve the Corev1 Client via clientset and list all Nodes in the cluster
	pods, err := cl.CoreV1().Pods(ns).List(context.TODO(), metav1.ListOptions{})
	// if ns != "" {
	// 	log.Printf("There are %d pods in the %v namespace\n", len(pods.Items), ns)
	// } else {
	// 	log.Printf("There are %d pods in the cluster\n", len(pods.Items))
	// }

	if err != nil {
		log.Fatal(err)
	}

	return pods

}
