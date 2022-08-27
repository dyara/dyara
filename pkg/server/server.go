package server

import (
	"github.com/dyara/dyara/pkg/server/middlewares"
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
}

func New() (*Server, error) {
	s := &Server{
		router: gin.Default(),
	}

	s.router.Use(middlewares.SetupCheck)

	return s, nil
}

func (server Server) Run(addr ...string) error {
	return server.router.Run(addr...)
}
