package main

import (
	"log"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func createKubeClient() (cl *kubernetes.Clientset, err error) {

	// Use the K8s ServiceAccount as kubeconfig
	log.Println("Using incluster K8S client")
	config, err := rest.InClusterConfig()

	if err != nil {
		log.Fatal(err)
	}

	// Create a clientset for the config
	clientset, err := kubernetes.NewForConfig(config)

	if err != nil {
		log.Fatal(err)
	}

	return clientset, err

}
