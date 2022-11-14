package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/Drozd0f/ttto-go/monolith/db"
	"github.com/Drozd0f/ttto-go/monolith/models"
	"github.com/google/uuid"
)

var ErrGameNotFound = errors.New("game not found")

func (r *Repository) CreateGame(ctx context.Context, ownerID uuid.UUID) (db.CreateGameRow, error) {
	g, err := r.q.CreateGame(ctx, db.CreateGameParams{
		ID:      uuid.New(),
		OwnerID: ownerID,
	})
	if err != nil {
		return db.CreateGameRow{}, fmt.Errorf("create game: %w", err)
	}

	return g, nil
}

func (r *Repository) GetTotalGames(ctx context.Context) (int64, error) {
	c, err := r.q.GetTotalGames(ctx)
	if err != nil {
		return 0, fmt.Errorf("get total games: %w", err)
	}

	return c, nil
}

func (r *Repository) GetGameByID(ctx context.Context, gID uuid.UUID) (db.GamesWithUsername, error) {
	g, err := r.q.GetGameById(ctx, gID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return db.GamesWithUsername{}, ErrGameNotFound
		}

		return db.GamesWithUsername{}, fmt.Errorf("queries get user by id: %w", err)
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

func (r *Repository) UpdateGame(ctx context.Context, g *models.Game) error {
	err := r.q.UpdateGameById(ctx, db.UpdateGameByIdParams{
		ID:              g.ID,
		OwnerID:         g.Owner.ID,
		OpponentID:      g.Opponent.GetID(),
		CurrentPlayerID: g.CurrentPlayer.GetID(),
		StepCount:       g.StepCount,
		WinnerID:        g.Winner.GetID(),
		Field:           g.MarshalField(),
		CurrentState:    g.CurrentState,
	})
	if err != nil {
		log.Println(err)
		return fmt.Errorf("update game by id: %w", err)
	}

	return nil
}
