package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/Drozd0f/ttto-go/services/auth/repository"
	"github.com/Drozd0f/ttto-go/services/auth/schemes"
)

var (
	ErrUserAlreadyExists = errors.New("user already exist")
)

func (s *Service) CreateUser(ctx context.Context, u *schemes.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.EncryptPassword(); err != nil {
		return fmt.Errorf("user encrypt password: %w", err)
	}

	if err := s.rep.CreateUser(ctx, u); err != nil {
		if errors.Is(err, repository.ErrUniqueConstraint) {
			return ErrUserAlreadyExists
		}
		return fmt.Errorf("repository create user: %w", err)
	}

	return nil
}
