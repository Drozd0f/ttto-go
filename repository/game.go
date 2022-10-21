package repository

import (
	"context"
	"fmt"

	"github.com/Drozd0f/ttto-go/db"
	"github.com/Drozd0f/ttto-go/models"
	"github.com/google/uuid"
)

func (r *Repository) CreateGame(ctx context.Context, ownerID uuid.UUID) (db.CreateGameRow, error) {
	g, err := r.q.CreateGame(ctx, ownerID)
	if err != nil {
		return db.CreateGameRow{}, fmt.Errorf("create game: %w", err)
	}

	return g, nil
}

func (r *Repository) GetGames(ctx context.Context, p models.Paginator) ([]db.GamesWithUsername, error) {
	games, err := r.q.GetGames(ctx, db.GetGamesParams{
		Offset: p.Offset(),
		Limit:  p.Limit,
	})
	if err != nil {
		return []db.GamesWithUsername{}, err
	}

	return games, nil
}
