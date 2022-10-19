package repository

import (
	"context"

	"github.com/Drozd0f/ttto-go/db"
	"github.com/Drozd0f/ttto-go/models"
)

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
