package main

import (
	"fmt"
	"net/http"
)

type Project struct {
	ID    string `json:"id"`
	PodID string `json:"value"`
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the RCK!"))
}

func resultHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	ns := "default"

	clientset, err := createKubeClient()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	pods := getPods(clientset, ns)
	// print pods
	for i, pod := range pods.Items {
		fmt.Printf("[%d] %s\n", i, pod.GetName())
		w.Write([]byte(pod.GetName() + "\n"))
	}

}

func returnHealth(w http.ResponseWriter, r *http.Request) {
	health := "Healthy"
	_, err := w.Write([]byte(health + "\n"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
