package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Project struct {
	ID      string `json:"id"`
	PodName string `json:"value"`
}

type Health struct {
	Status string `json:"Status"`
}

// Home Page Handler
func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the RCK!"))
}

func projectsHandler(w http.ResponseWriter, r *http.Request) {
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
		w.Write([]byte("Pod Name: " + pod.GetName() + "  Namespace: " + ns + "\n"))
	}

}

// Return Healthy (future Prometheus integration)
func returnHealth(w http.ResponseWriter, r *http.Request) {
	// health := "Healthy"
	health := Health{Status: "Healthy"}

	jsonResponse, err := json.Marshal(health)
	if err != nil {
		fmt.Println("Unable to encode JSON")
	}

	// fmt.Println(string(jsonResponse))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

// Return Healthy (future Prometheus integration)
func returnUnhealth(w http.ResponseWriter, r *http.Request) {
	// health := "Unhealth"
	health := Health{Status: "Unhealth"}

	jsonResponse, err := json.Marshal(health)
	if err != nil {
		fmt.Println("Unable to encode JSON")
	}

	// fmt.Println(string(jsonResponse))
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

// Return the Hostname of the node where is running
func returnHostname(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	w.WriteHeader(http.StatusOK)

	hostname, err := os.Hostname()
	if err != nil {
		hostname = "Unknown"
	}

	hostnameStr := "Hostname: " + hostname

	// return the hostname in json format
	if err := json.NewEncoder(w).Encode(hostnameStr); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
