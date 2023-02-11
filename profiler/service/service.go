package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/Drozd0f/ttto-go/profiler/config"
	"github.com/Drozd0f/ttto-go/profiler/repository"
	"github.com/Drozd0f/ttto-go/schemes"
)

var ErrUserAlreadyExists = errors.New("user already exist")

type Service struct {
	repository *repository.Repository
	c          *config.Config
}

func New(r *repository.Repository, c *config.Config) *Service {
	return &Service{
		repository: r,
		c:          c,
	}
}

func (s *Service) CreateUser(ctx context.Context, u *schemes.User) (*schemes.User, error) {
	if err := u.Validate(); err != nil {
		return nil, err
	}

	if err := u.EncryptPassword(); err != nil {
		return nil, fmt.Errorf("user encrypt password: %w", err)
	}

	u, err := s.repository.CreateUser(ctx, u)
	if err != nil {
		if errors.Is(err, repository.ErrUniqueConstraint) {
			return nil, ErrUserAlreadyExists
		}

		return nil, fmt.Errorf("repository create user: %w", err)
	}

	return u, nil
}

func (s *Service) GetUserByUsername(ctx context.Context, u *schemes.User) (*schemes.User, error) {
	u, err := s.repository.GetUserByUsername(ctx, u)
	if err != nil {
		return nil, fmt.Errorf("repository get user: %w", err)
	}

	return u, nil
}
