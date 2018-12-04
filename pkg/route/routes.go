package route

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Johnlovescoding/ENPM613/HOLMS/pkg/apiserver"
	"github.com/Johnlovescoding/ENPM613/HOLMS/pkg/authserver"
	"github.com/gorilla/mux"
)

type Route struct {
	path     string
	funcName func(w http.ResponseWriter, r *http.Request)
	method   string
}

func AddRoutes(router *mux.Router) {
	Routes := []Route{
		{"/student", auth(apiserver.GetStudent), "GET"},
		{"/student", auth(apiserver.PatchStudent), "PATCH"},
		//{"/student", auth(apiserver.RegisterCourse), "PATCH"},
		{"/student", auth(apiserver.PostStudent), "POST"},
		{"/student/all", auth(apiserver.ListAllStudent), "GET"},
		{"/course/all", auth(apiserver.Health), "OPTIONS"},
		{"/login", authserver.Login, "POST"},
		{"/signup", apiserver.PostStudent, "POST"},

		{"/account", auth(authserver.Account), "GET"},
		{"/discussionboard", apiserver.ListAllComment, "GET"},
		{"/discussionboard", apiserver.PostComment, "POST"},

		{"/course", apiserver.Health, "OPTIONS"},
		{"/course", apiserver.Health, "OPTIONS"},
		{"/student", apiserver.Health, "OPTIONS"},
		{"/student/all", apiserver.Health, "OPTIONS"},
		{"/login", apiserver.Health, "OPTIONS"},
		{"/account", apiserver.Health, "OPTIONS"},
		{"/signup", apiserver.Health, "OPTIONS"},

		{"/discussionboard", apiserver.Health, "OPTIONS"},
		{"/discussionboard", apiserver.Health, "OPTIONS"},
	}

	for _, route := range Routes {
		router.HandleFunc(route.path, route.funcName).Methods(route.method)
	}

}

func auth(f func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authorized, msg, code := authserver.Authorize(w, r)
		if authorized {
			f(w, r)

		} else {

			response, err := json.Marshal(msg)
			if err != nil {
				log.Println(err)
			}
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, PATCH")
			w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
			w.WriteHeader(code)
			w.Write(response)
		}
	})
}
