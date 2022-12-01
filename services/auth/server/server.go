package server

import (
	"github.com/Drozd0f/ttto-go/gen/proto/auth"
	"github.com/Drozd0f/ttto-go/services/auth/conf"
	"github.com/Drozd0f/ttto-go/services/auth/service"
)

type Server struct {
	auth.UnimplementedAuthServer
	c          *conf.Config
	service *service.Service
}

var _ auth.AuthServer = (*Server)(nil)

func New(c *conf.Config, service *service.Service) *Server {
	return &Server{
		c:                       c,
		service:                 service,
	}
}
