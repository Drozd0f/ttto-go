package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/Drozd0f/ttto-go/services/auth/repository"
	"github.com/Drozd0f/ttto-go/services/auth/schemes"
	"github.com/golang-jwt/jwt/v4"
)

var (
	ErrUserAlreadyExists = errors.New("user already exist")
	ErrUserNotExists     = errors.New("user not exist")
	ErrUserInvalidData   = errors.New("invalid data")
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

func (s *Service) LoginUser(ctx context.Context, u *schemes.User) (string, error) {
	if err := u.Validate(); err != nil {
		return "", err
	}

	storUser, err := s.rep.GetUserByName(ctx, u)
	if err != nil {
		if errors.Is(err, repository.ErrNoResult) {
			return "", ErrUserNotExists
		}

		return "", fmt.Errorf("repository get user: %w", err)
	}

	if err = u.CheckPassword(storUser.Password); err != nil {
		if errors.Is(err, schemes.ErrInvalidData) {
			return "", ErrUserInvalidData
		}

		return "", fmt.Errorf("user check password: %w", err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"ID":       storUser.ID,
		"Username": storUser.Username,
	})

	tokenString, err := token.SignedString([]byte(s.cfg.Secret))
	if err != nil {
		return "", fmt.Errorf("generate token: %w", err)
	}

	return tokenString, nil
}
