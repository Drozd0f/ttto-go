package service

import (
	"context"
	"fmt"
	"net/url"

	"github.com/Drozd0f/ttto-go/models"
)

// func (s *Service) CreateGame(ctx context.Context) (models.GameWithUsername, error) {
// 	g, err := s.r.CreateGame(ctx)
// 	if err != nil {
// 		return models.GameWithUsername{}, fmt.Errorf("repository create game: %w", err)
// 	}

// 	return *models.NewGameWithUsernameFromDB(g), nil
// }

func (s *Service) GetGames(ctx context.Context, v url.Values) ([]models.GameWithUsername, error) {
	storGames, err := s.r.GetGames(ctx, models.NewPaginatorFromQuery(v))
	if err != nil {
		return nil, fmt.Errorf("repository get games: %w", err)
	}

	return models.GameWithUsernameSliceFromDB(storGames), nil
}
