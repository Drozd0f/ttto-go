package service

import (
	"context"
	"errors"
	"fmt"
	"net/url"

	"github.com/Drozd0f/ttto-go/models"
	"github.com/Drozd0f/ttto-go/repository"
	"github.com/google/uuid"
)

var (
	ErrGameNotFound = errors.New("game not found")
	ErrInvalidState = errors.New("invalid state")
	ErrUserInGame   = errors.New("user already in game")
)

func (s *Service) CreateGame(ctx context.Context) (models.Game, error) {
	u := ctx.Value("user").(models.User)
	g, err := s.r.CreateGame(ctx, u.ID)
	if err != nil {
		return models.Game{}, fmt.Errorf("repository create game: %w", err)
	}

	return *models.NewGameFromDB(g), nil
}

func (s *Service) GetGameByID(ctx context.Context, gameID string) (*models.Game, error) {
	gID, err := uuid.Parse(gameID)
	if err != nil {
		return nil, ErrInvalidId
	}

	g, err := s.r.GetGameByID(ctx, gID)
	if err != nil {
		if errors.Is(err, repository.ErrGameNotFound) {
			return nil, ErrGameNotFound
		}

		return nil, fmt.Errorf("repository get user: %w", err)
	}

	return models.NewGameFromDB(g), nil
}

func (s *Service) GetGames(ctx context.Context, v url.Values) (models.PaginationGameSlice, error) {
	p := models.NewPaginatorFromQuery(v)
	storGames, err := s.r.GetGames(ctx, p)
	if err != nil {
		return models.PaginationGameSlice{}, fmt.Errorf("repository get games: %w", err)
	}

	countGames, err := s.r.GetTotalGames(ctx)
	if err != nil {
		return models.PaginationGameSlice{}, fmt.Errorf("repository get total games: %w", err)
	}

	p.SetTotalPages(countGames)

	return models.PaginationGameSlice{
		Games:     models.NewGameSliceFromDB(storGames),
		Paginator: p,
	}, nil
}

func (s *Service) LoginGame(ctx context.Context, gameID string) (*models.Game, error) {
	u := ctx.Value("user").(models.User)
	g, err := s.GetGameByID(ctx, gameID)
	if err != nil {
		return nil, err
	}

	if g.CurrentState != models.StatePending {
		return nil, ErrInvalidState
	}

	if g.Owner.ID == u.ID {
		return nil, ErrUserInGame
	}

	g.SetOpponent(u)

	if err = s.r.UpdateGame(ctx, g); err != nil {
		return nil, fmt.Errorf("repository update game: %w", err)
	}

	return g, nil
}
