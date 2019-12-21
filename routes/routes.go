package routes

import (
	"net/http"

	"github.com/isgo-golgo13/gorilla-restsvc/route_handlers"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

//Add route handlers here
var Route_Entries = Routes{
	Route{
		"EnginesIndex",
		"GET",
		"/engines",
		route_handlers.GetEngines,
	},
}
