package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"

	"github.com/Drozd0f/ttto-go/models"
	"github.com/Drozd0f/ttto-go/repository"
)

var ErrUserNotFound = errors.New("user not found")

func (s *Service) GetUserByID(ctx context.Context, userId string) (models.User, error) {
	uId, err := uuid.Parse(userId)
	if err != nil {
		return models.User{}, ErrInvalidId
	}

	u, err := s.r.GetUserByID(ctx, uId)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			return models.User{}, ErrUserNotFound
		}

		return models.User{}, fmt.Errorf("repository get user: %w", err)
	}

	return models.User{
		Username: u.Username,
		Password: u.Password,
	}, nil
}
