package service

import (
	"github.com/Drozd0f/ttto-go/monolith/conf"
	"github.com/Drozd0f/ttto-go/monolith/models"
	"github.com/Drozd0f/ttto-go/monolith/repository"
)

type Service struct {
	r *repository.Repository
	c *conf.Config
	l *models.Lobby
}

func New(r *repository.Repository, c *conf.Config) *Service {
	return &Service{
		r: r,
		c: c,
		l: models.NewLobby(),
	}
}
