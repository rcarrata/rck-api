package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

// Func for create k8s clients inside of a k8s container
func createKubeClient() (cl *kubernetes.Clientset, err error) {

	// Use the K8s ServiceAccount as kubeconfig
	log.Println("Using incluster K8S client")
	config, err := rest.InClusterConfig()
	if err != nil {
		log.Fatal(err)
	}

	// Create a new clientset for the config
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}

	return clientset, err

}

// Marshal the json content and send it as response of the API
func sendJsonResponse(jsoncontent Health, w http.ResponseWriter) {
	jsonResponse, err := json.Marshal(jsoncontent)
	if err != nil {
		fmt.Println("Unable to encode JSON")
	}

	// fmt.Println(string(jsonResponse))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)

}
