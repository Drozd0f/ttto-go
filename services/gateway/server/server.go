package server

import (
	"github.com/Drozd0f/ttto-go/services/gateway/server/middleware"
	"github.com/gin-gonic/gin"

	"github.com/Drozd0f/ttto-go/services/gateway/conf"
	"github.com/Drozd0f/ttto-go/services/gateway/service"
)

type Server struct {
	*gin.Engine
	c       *conf.Config
	service *service.Service
}

func New(c *conf.Config, service *service.Service) (*Server, error) {
	if !c.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	s := &Server{
		Engine:  gin.Default(),
		c:       c,
		service: service,
	}

	s.registerMiddlewares()
	s.registerHandlers()

	return s, nil
}

func (s *Server) registerMiddlewares() {
	s.Use(middleware.LoggingHandler)
	s.Use(middleware.ErrorHandler)
}

func (s *Server) registerHandlers() {
	v2 := s.Group("api/v2")
	v2.GET("/ping", s.ping)

	s.registerAuthHandlers(v2)
}
