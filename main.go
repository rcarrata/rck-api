package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/rcarrata/rck/pkg/api"
)

func main() {

	// Import from config.json file the listening Port Mux Router
	port := ":" + os.Getenv("PORT")

	// Configure Mux Router and Initialize the routes
	router := ConfigureRouter()

	// Run the Mux Router with the specific config
	log.Printf("Server running in port %s", port)
	log.Fatal(http.ListenAndServe(port, handlers.LoggingHandler(os.Stdout, router)))

}
