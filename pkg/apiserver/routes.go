package apiserver

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	path     string
	funcName func(w http.ResponseWriter, r *http.Request)
	method   string
}

func AddRoutes(router *mux.Router) {
	Routes := []Route{
		{"/student", GetStudent, "GET"},
		{"/student", PostStudent, "POST"},
		{"/student/all", ListAllStudent, "GET"},
		{"/course/all", Health, "OPTIONS"},
		
		{"/course", Health, "OPTIONS"},
		{"/course", Health, "OPTIONS"},
		{"/student", Health, "OPTIONS"},
		{"/student/all", Health, "OPTIONS"},
	}

	for _, route := range Routes {
		router.HandleFunc(route.path, route.funcName).Methods(route.method)
	}

}
