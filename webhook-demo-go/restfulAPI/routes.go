package restfulAPI

import (
	"net/http"
	"utils"

	"github.com/gorilla/mux"
)

type Route struct {
	Name    string
	Method  string
	Pattern string
	handler http.HandlerFunc
}

type Routes []Route

var (
	routes = Routes{
		Route{
			Name:    "UserGuide",
			Method:  "GET",
			Pattern: "/",
			handler: UserGuide,
		},
		Route{
			Name:    "GetCurrentWarn",
			Method:  "GET",
			Pattern: "/get",
			handler: GetCurrentWarn,
		},
		Route{
			Name:    "PostMonitorWarn",
			Method:  "POST",
			Pattern: "/add",
			handler: PostMonitorWarn,
		},
	}
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		handler := utils.Logger(route.handler, route.Name)
		router.
			Methods(route.Method).
			Name(route.Name).
			Path(route.Pattern).
			Handler(handler)
	}

	return router
}
