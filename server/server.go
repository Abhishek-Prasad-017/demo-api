package server

import (
	"github.com/buaazp/fasthttprouter"
	"github.com/demo-api/service"
	"github.com/spf13/viper"
	"github.com/valyala/fasthttp"
	"log"
)

//Server struct
type Server struct {
	router       *fasthttprouter.Router
	mongoService *service.MongoService

}


func StartServer(path string, viper *viper.Viper) error {
	//Create mongodb service
	log.Println("I am in server.go")
	mongo, err := service.LoadMongo(viper)
	if err != nil {
		log.Fatalf("Could not connect to Mongo %s", err)
	}

	s := &Server{
		router:       fasthttprouter.New(),
		mongoService: mongo,
	}

	//Generate API pathsserve
	s.setupRoutes()

	//Start the server for listening
	log.Println("Server Listening........")
	return fasthttp.ListenAndServe(path, s.router.Handler)

}
