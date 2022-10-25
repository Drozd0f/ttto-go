package server

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"

	"github.com/Drozd0f/ttto-go/conf"
	"github.com/Drozd0f/ttto-go/server/middleware"
	"github.com/Drozd0f/ttto-go/service"
)

// serv -> service -> db
type Server struct {
	*gin.Engine
	upgrader websocket.Upgrader
	service  *service.Service
	c        *conf.Config
}

func New(s *service.Service, c *conf.Config) *Server {
	if !c.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	serv := &Server{
		Engine:   gin.Default(),
		upgrader: websocket.Upgrader{},
		service:  s,
		c:        c,
	}
	serv.RegisterHandlers()

	return serv
}

func (s *Server) RegisterHandlers() {
	v1 := s.Group("/api/v1", middleware.Auth(s.c.Secret))
	v1.GET("/ping", s.ping)

	s.registerAuthHandlers(v1)
	s.registerGameHandlers(v1)
}
