package routes

import (
	"github.com/edigar/socialnets-web/src/middlewares"
	"github.com/gorilla/mux"
	"net/http"
)

type Route struct {
	URI                    string
	Method                 string
	Function               func(w http.ResponseWriter, r *http.Request)
	AuthenticationRequired bool
}

func Setup(r *mux.Router) *mux.Router {
	routes := loginRoute
	routes = append(routes, userRoutes...)
	routes = append(routes, postRoutes...)
	routes = append(routes, homeRoute)
	routes = append(routes, logoutRoute)

	for _, route := range routes {
		if route.AuthenticationRequired {
			r.HandleFunc(route.URI, middlewares.Logger(middlewares.Authenticate(route.Function))).Methods(route.Method)
		} else {
			r.HandleFunc(route.URI, middlewares.Logger(route.Function)).Methods(route.Method)
		}
	}

	for _, route := range routes {
		r.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}

	fileServer := http.FileServer(http.Dir("./assets/"))
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

	return r
}
