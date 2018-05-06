package controller

import (
	"net/http"

	"github.com/gorilla/mux"
)

// StartWebServer starts the http server for this service
// on the given http port.
func StartWebServer(port string) {
	http.Handle("/", NewRouter())

	log.Info("Starting HTTP service at ", port)
	err := http.ListenAndServe(":"+port, nil)

	if err != nil {
		log.Panicf("An error occured starting HTTP listener at port %s, error %s", port, err)
	}
}

// NewRouter returns a pointer to a mux.Router we can use as a handler.
func NewRouter() *mux.Router {
	// Create an instance of the Gorilla router
	// Gorilla router matches incoming requests against a list of
	// registered routes and calls a handler for the route that matches
	// the URL or other conditions
	router := mux.NewRouter().StrictSlash(true)

	// Get for query graphql
	router.Methods("GET").
		Path("/graphql").
		Name("GetGraphql").
		HandlerFunc(httpGet) // TODO put graphql function

	// Post for graphql to create users
	router.Methods("POST").
		Path("/graphql").
		Name("PostGraphql").
		HandlerFunc(httpPost) // TODO put graphql function

	// get health status of this service.
	router.Methods("GET").
		Path("/health").
		Name("health").
		HandlerFunc(Health) // what's the health

	return router
}
