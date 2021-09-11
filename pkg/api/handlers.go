package api

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gorilla/mux"
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

// Get objects from the project default
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

	w.Write([]byte("## Checking Namespace -> " + ns + "\n"))
	for _, pod := range pods.Items {
		// fmt.Printf("[%d] %s\n", i, pod.GetName())
		w.Write([]byte("Pod Name: " + pod.GetName() + "\n"))
	}

}

// Get objects from the Projects Ids from Path
func getProjectHandler(w http.ResponseWriter, r *http.Request) {

	// Return the Id in the request of the path
	vars := mux.Vars(r)

	// Asign the NS from the Id
	ns := vars["Id"]
	// fmt.Fprintln(w, "Project Id:", ns)

	clientset, err := createKubeClient()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	pods := getPods(clientset, ns)
	// print pods

	w.WriteHeader(http.StatusOK)

	w.Write([]byte("## Checking Namespace -> " + ns + "\n"))
	for _, pod := range pods.Items {
		// fmt.Printf("[%d] %s\n", i, pod.GetName())
		w.Write([]byte("Pod Name: " + pod.GetName() + "\n"))
	}

}

// Return Healthy (future Prometheus integration)
func returnHealth(w http.ResponseWriter, r *http.Request) {
	// health := "Healthy"
	health := Health{Status: "Healthy"}

	sendJsonResponse(health, w)
}

// Return Unhealthy (future Prometheus integration)
func returnUnhealth(w http.ResponseWriter, r *http.Request) {
	// health := "Unhealth"
	health := Health{Status: "Unhealth"}

	sendJsonResponse(health, w)
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
