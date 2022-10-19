package service

import (
	"context"
	"fmt"

	"github.com/Drozd0f/ttto-go/models"
)

func (s *Service) Reg(ctx context.Context, u *models.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	if err := s.r.CreateUser(ctx, u); err != nil {
		return fmt.Errorf("repository create user: %w", err)
	}

	return nil
}

func (s *Service) Login(ctx context.Context, u *models.User) (string, error) {
	if err := u.Validate(); err != nil {
		return "", err
	}

	storUser, err := s.r.GetUserByName(ctx, u)
	if err != nil {
		return "", fmt.Errorf("repository get user: %w", err)
	}

	fmt.Println(storUser) // t, err := tokenGenerate(storUser)
	return "Token", nil
}
