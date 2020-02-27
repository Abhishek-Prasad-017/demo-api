package server

import (
	"github.com/valyala/fasthttp"
	"log"
)

type Routes []Route

type Route struct {
	Method      string
	Path        string
	HandlerFunc fasthttp.RequestHandler
}

func (s *Server) setupRoutes() {
	var routes = Routes{

		Route{

			"GET",
			"/",
			s.GetSoftwareDetails,
		},

	}


	for _, route := range routes {
		s.router.Handle(route.Method, route.Path, fasthttp.CompressHandler(route.HandlerFunc))
	}
	log.Println("I am in routes.go. Routes set up.")
}
