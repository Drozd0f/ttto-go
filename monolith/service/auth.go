package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v4"

	"github.com/Drozd0f/ttto-go/monolith/models"
	"github.com/Drozd0f/ttto-go/monolith/repository"
)

var (
	ErrUserNotExists     = errors.New("user not exist")
	ErrUserAlreadyExists = errors.New("user already exist")
	ErrUserInvalidData   = errors.New("invalid data")
)

func (s *Service) Reg(ctx context.Context, u *models.User) error {
	if err := u.Validate(); err != nil {
		return err
	}

	if err := u.EncryptPassword(); err != nil {
		return fmt.Errorf("user encrypt password: %w", err)
	}

	if err := s.r.CreateUser(ctx, u); err != nil {
		if errors.Is(err, repository.ErrUniqueConstraint) {
			return ErrUserAlreadyExists
		}
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
		if errors.Is(err, repository.ErrNoResult) {
			return "", ErrUserNotExists
		}
		return "", fmt.Errorf("repository get user: %w", err)
	}

	if err = u.CheckPassword(storUser.Password); err != nil {
		if errors.Is(err, models.ErrInvalidData) {
			return "", ErrUserInvalidData
		}
		return "", fmt.Errorf("user check password: %w", err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"ID":       storUser.ID,
		"Username": storUser.Username,
	})

	tokenString, err := token.SignedString([]byte(s.c.Secret))
	if err != nil {
		return "", fmt.Errorf("generate token: %w", err)
	}

	return tokenString, nil
}
