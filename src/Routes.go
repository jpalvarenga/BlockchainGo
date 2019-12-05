package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route struct represents a route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// InitRouter initializes mux routers in routes
func InitRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}

var routes = []Route{
	Route{
		"UploadBlockchain",
		"GET",
		"/upload",
		UploadBlockchain,
	},
	Route{
		"UploadBlock",
		"GET",
		"/block/{height}/{hash}",
		UploadBlock,
	},
	Route{
		"ReceiveBlock",
		"POST",
		"/block/receive",
		ReceiveBlock,
	},
	Route{
		"Start",
		"GET",
		"/start",
		Start,
	},
	Route{
		"RegisterPeer",
		"GET",
		"/peer",
		RegisterPeer,
	},
}
