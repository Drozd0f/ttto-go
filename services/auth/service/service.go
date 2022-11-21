package service

import (
	"github.com/Drozd0f/ttto-go/services/auth/conf"
	"github.com/Drozd0f/ttto-go/services/auth/repository"
)

type Service struct {
	rep *repository.Repository
	cfg *conf.Config
}

func New(rep *repository.Repository, c *conf.Config) *Service {
	return &Service{
		rep: rep,
		cfg: c,
	}
}
