package server

import (
	"github.com/gin-gonic/gin"

	"github.com/Drozd0f/ttto-go/service"
)

// serv -> service -> db
type Server struct {
	*gin.Engine
	service *service.Service
}

func New(s *service.Service) *Server {
	serv := &Server{
		Engine:  gin.Default(),
		service: s,
	}

	serv.RegisterHandlers()

	return serv
}

func (s *Server) RegisterHandlers() {
	v1 := s.Group("/api/v1")
	v1.GET("/ping", s.Ping)
	v1.GET("/users/:userId", s.GetUser)
	v1.POST("/reg", s.Reg)
}
