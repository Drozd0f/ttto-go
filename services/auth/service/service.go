package service

import (
	"github.com/Drozd0f/ttto-go/services/auth/repository"
)

type Service struct {
	rep *repository.Repository
}

func New(rep *repository.Repository) *Service {
	return &Service{rep: rep}
}
