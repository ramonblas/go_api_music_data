package controller

import (
	"net/http"

	"github.com/gorilla/mux"
	"go_api/src/service"
)

type Route struct {
	Name       string
	Method     string
	Pattern    string
	HandleFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		router.Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandleFunc)
	}

	return router
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		actions.Index,
	},
	Route{
		"SongList",
		"GET",
		"/cancions",
		actions.SongList,
	},
	Route{
		"Songshow",
		"GET",
		"/cancion/{id}",
		actions.SongShow,
	},
	Route{
		"SongAdd",
		"POST",
		"/cancion",
		actions.SongAdd,
	},
	Route{
		"SongUpdate",
		"PUT",
		"/cancion/{id}",
		actions.SongUpdate,
	},
	Route{
		"SongRemove",
		"DELETE",
		"/cancion/{id}",
		actions.SongRemove,
	},
}
