package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/url"

	"github.com/Drozd0f/ttto-go/models"
)

func (s *Service) GetGames(ctx context.Context, v url.Values) ([]models.Game, error) {
	storGames, err := s.r.GetGames(ctx, models.NewPaginatorFromQuery(v))
	if err != nil {
		return nil, fmt.Errorf("repository get games: %w", err)
	}

	games := make([]models.Game, 0, len(storGames))
	for _, g := range storGames {
		var f models.Field
		if err := json.Unmarshal(g.Field.RawMessage, &f); err != nil {
			log.Println(err.Error())
		}

		games = append(games, models.Game{
			ID:              g.ID,
			OwnerID:         g.OwnerID,
			OpponentID:      g.OpponentID,
			CurrentPlayerID: g.CurrentPlayerID,
			StepCount:       g.StepCount,
			WinnerID:        g.WinnerID,
			Field:           f,
			CurrentState:    g.CurrentState,
		})
	}

	return games, nil
}
