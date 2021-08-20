package main

import (
	"github.com/gorilla/mux"
)

//ConfigureRouter setup the router
func ConfigureRouter() *mux.Router {
	router := mux.NewRouter()

	// router.PathPrefix("/static").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))

	router.HandleFunc("/", homeHandler)
	router.HandleFunc("/projects", projectsHandler)
	router.HandleFunc("/health", returnHealth)
	router.HandleFunc("/host", returnHostname)
	router.HandleFunc("/test", test)

	return router
}
