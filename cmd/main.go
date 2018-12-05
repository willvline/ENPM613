package main

import (
	"log"
	"net/http"

	config "github.com/Johnlovescoding/ENPM613/HOLMS/pkg/config"
	mongo "github.com/Johnlovescoding/ENPM613/HOLMS/pkg/mongo"
	"github.com/Johnlovescoding/ENPM613/HOLMS/pkg/route"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var conf = config.Config{}
var mongoDB = mongo.MongoDB{}

// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {
	//conf.Read()

	//mongoDB.Server = conf.Server
	//mongoDB.Database = conf.Database
	mongoDB.Server = "mongo:27017"
	mongoDB.Database = "HOLMS_db"
	mongoDB.Connect()
}

func main() {

	router := mux.NewRouter()
	router.Use(handlers.ProxyHeaders)
	route.AddRoutes(router)
	err := http.ListenAndServe("0.0.0.0:8000", router)
	if err != nil {
		log.Println(err)
	}
}
