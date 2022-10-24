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
	ErrNotYourTurn  = errors.New("not your turn")
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

func (s *Service) getGameWithState(ctx context.Context, gameID string, state int16) (*models.Game, error) {
	g, err := s.GetGameByID(ctx, gameID)
	if err != nil {
		return nil, err
	}

	if g.CurrentState != state {
		return nil, ErrInvalidState
	}

	return g, nil
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

func (s *Service) UpdateGame(ctx context.Context, g *models.Game) error {
	if err := s.r.UpdateGame(ctx, g); err != nil {
		return fmt.Errorf("repository update game: %w", err)
	}

	return nil
}

func (s *Service) LoginGame(ctx context.Context, gameID string) (*models.Game, error) {
	g, err := s.getGameWithState(ctx, gameID, models.StatePending)
	if err != nil {
		return nil, err
	}

	u := ctx.Value("user").(models.User)
	if g.Owner.ID == u.ID {
		return nil, ErrUserInGame
	}

	g.SetOpponent(u)

	if err = s.UpdateGame(ctx, g); err != nil {
		return nil, err
	}

	return g, nil
}

func (s *Service) MakeStep(ctx context.Context, gameID string, c models.Coord) (*models.Game, error) {
	if err := c.Validate(); err != nil {
		return nil, err
	}

	g, err := s.getGameWithState(ctx, gameID, models.StateInGame)
	if err != nil {
		return nil, err
	}

	u := ctx.Value("user").(models.User)
	if g.CurrentPlayer.Player.ID != u.ID {
		return nil, ErrNotYourTurn
	}

	if err = g.MakeStep(c); err != nil {
		return nil, err
	}

	if err = s.UpdateGame(ctx, g); err != nil {
		return nil, err
	}

	return g, nil
}
