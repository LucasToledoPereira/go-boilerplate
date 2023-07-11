package server

import (
	"fmt"

	"github.com/LucasToledoPereira/go-boilerplate/config"
	"github.com/LucasToledoPereira/go-boilerplate/internal/router"
)

type Server struct {
	router *router.Router
}

func New(r *router.Router) (server *Server) {
	return &Server{
		router: r,
	}
}

func (s *Server) Run() {
	fmt.Printf("collection manager listening on port '%s'", config.C.Server.Port)
	fmt.Println()

	panic(s.router.Router.Run(":" + config.C.Server.Port))
}
