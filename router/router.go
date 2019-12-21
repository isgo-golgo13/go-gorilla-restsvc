package router

import (
	"github.com/gorilla/mux"
	"net/http"
	"github.com/isgo-golgo13/gorilla-restsvc/routes"
	"github.com/isgo-golgo13/gorilla-restsvc/logger"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes.Route_Entries {
		var handler http.Handler
		handler = route.HandlerFunc = route.HandlerFunc
		handler = logger.Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	return router
}
