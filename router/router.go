package router

import (
	"github.com/gorilla/mux"
)

// NewRouter builds and returns a new router from routes
func NewRouter() *mux.Router {
	// When StrictSlash == true, if the route path is "/path/", accessing "/path" will perform a redirect to the former and vice versa.
	router := mux.NewRouter().StrictSlash(true)

	//Setting the same matching conditions again and again can be boring, so we have a way to group several routes that share the same requirements.
	sub := router.PathPrefix("/api/v1").Subrouter()

	// Append routes for all the objects
	routes := myRoutes

	for _, route := range routes {
		sub.
			HandleFunc(route.Pattern, route.HandlerFunc).
			Name(route.Name).
			Methods(route.Method)
	}

	return router
}
