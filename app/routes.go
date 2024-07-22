package app

import "github.com/kusuma-lint/website-go/app/controllers"

func (server *Server) initializeRoutes() {
	server.Router.HandleFunc("/", controllers.Home).Methods("GET")
	// server.Router.HandleFunc("/", controllers.Home).Methods("GET")
}
