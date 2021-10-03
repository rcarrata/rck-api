package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	v1 "k8s.io/api/core/v1"
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

func (project *Projects) AddItem(item Project) {
	project.Items = append(project.Items, item)
}

func aggregateProject(pods *v1.PodList, svcs *v1.ServiceList, ns string, w http.ResponseWriter) {

	podsList := Projects{}
	svcsList := Projects{}

	for _, pod := range pods.Items {
		podsStat := Project{PodName: pod.GetName()}
		podsList.AddItem(podsStat)
	}

	for _, svc := range svcs.Items {
		svcStat := Project{SVCName: svc.GetName()}
		svcsList.AddItem(svcStat)
	}

	w.Write([]byte("## Checking Namespace -> " + ns + "\n"))
	w.Write([]byte("#### List of Pods in Namespace -> " + ns + "\n"))
	for _, pod := range podsList.Items {
		// fmt.Printf("[%d] %s\n", i, pod.GetName())
		w.Write([]byte("Pod Name: " + pod.PodName + "\n"))

	}
	w.Write([]byte("\n"))
	w.Write([]byte("#### List of SVCs in Namespace -> " + ns + "\n"))
	for _, pod := range svcsList.Items {
		// fmt.Printf("[%d] %s\n", i, pod.GetName())
		w.Write([]byte("SVC Name: " + pod.SVCName + "\n"))
	}

}

// Not using Viper in this version
// func loadApiConfig(key string) string {

// 	// Set the location and the name of the config file
// 	viper.SetConfigName("app")
// 	viper.SetConfigType("env")
// 	viper.AddConfigPath("config")

// 	err := viper.ReadInConfig()
// 	if err != nil {
// 		log.Fatalf("Error while reading config file %s", err)
// 	}

// 	config, ok := viper.Get(key).(string)

// 	if !ok {
// 		log.Fatalf("Invalid type assertion")
// 	}

// 	return config
// }
