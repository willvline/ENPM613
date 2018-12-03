package main

import (
	"log"
	"net/http"

	apiserver "github.com/ENPM613/HOLMS/pkg/apiserver"
	config "github.com/ENPM613/HOLMS/pkg/config"
	mongo "github.com/ENPM613/HOLMS/pkg/mongo"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var conf = config.Config{}
var mongoDB = mongo.MongoDB{}

// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {
	conf.Read()

	mongoDB.Server = conf.Server
	mongoDB.Database = conf.Database
	mongoDB.Connect()
}

func main() {

	router := mux.NewRouter()
	router.Use(handlers.ProxyHeaders)
	apiserver.AddRoutes(router)
	err := http.ListenAndServe("localhost:8000", router)
	if err != nil {
		log.Println(err)
	}
}
