package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

func main() {

	// TODO: convert to a config file and import
	port := ":8080"

	// Configure Mux Router and Initialize the routes
	router := ConfigureRouter()

	log.Printf("Server running in port %s", port)
	log.Fatal(http.ListenAndServe(port, handlers.LoggingHandler(os.Stdout, router)))

}
