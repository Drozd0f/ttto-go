package server

import (
	"github.com/gin-gonic/gin"

	"github.com/Drozd0f/ttto-go/services/gateway/conf"
	"github.com/Drozd0f/ttto-go/services/gateway/server/middleware"
	"github.com/Drozd0f/ttto-go/services/gateway/service"
)

type Server struct {
	*gin.Engine
	c          *conf.Config
	service *service.Service
}

func New(c *conf.Config, service *service.Service) (*Server, error) {
	if !c.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	s := &Server{
		Engine:     gin.Default(),
		c:          c,
		service: service,
	}
	if err := s.Engine.SetTrustedProxies(nil); err != nil {
		return nil, err
	}

	s.RegisterHandlers()

	return s, nil
}

func (s *Server) RegisterHandlers() {
	s.Use(middleware.ErrorHandler)

	v2 := s.Group("api/v2")
	v2.GET("/ping", s.ping)

	s.registerAuthHandlers(v2)
}
