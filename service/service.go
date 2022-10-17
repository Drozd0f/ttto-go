package service

import (
	"errors"

	"github.com/Drozd0f/ttto-go/repository"
)

var ErrInvalidId = errors.New("invalid id")

type Service struct {
	r *repository.Repository
}

func New(r *repository.Repository) *Service {
	return &Service{
		r: r,
	}
}
