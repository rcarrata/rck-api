package main

import (
	"log"

	"github.com/spf13/viper"
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

// Use Viper to Import the config file
func loadApiConfig(key string) string {

	// Set the location and the name of the config file
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	config, ok := viper.Get(key).(string)

	if !ok {
		log.Fatalf("Invalid type assertion")
	}

	return config
}
