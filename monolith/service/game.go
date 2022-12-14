package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/url"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"

	"github.com/Drozd0f/ttto-go/monolith/models"
	"github.com/Drozd0f/ttto-go/monolith/repository"
)

var (
	ErrGameInvalidId     = errors.New("invalid id")
	ErrGameNotFound      = errors.New("game not found")
	ErrGameInvalidState  = errors.New("invalid state")
	ErrGameUserExists    = errors.New("user already in game")
	ErrGameUserNotExists = errors.New("user not in game")
	ErrGameNotTurnUser   = errors.New("not your turn")
	ErrGameCellOccupied  = errors.New("cell is already occupied")
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
		return nil, ErrGameInvalidId
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

func (s *Service) UpdateGame(ctx context.Context, g *models.Game) error {
	if err := s.r.UpdateGame(ctx, g); err != nil {
		return fmt.Errorf("repository update game: %w", err)
	}

	s.l.Update(g) // update lobby by game id in all channel in this lobby

	return nil
}

func (s *Service) LoginGame(ctx context.Context, gameID string) (*models.Game, error) {
	g, err := s.getGameWithState(ctx, gameID, models.StatePending)
	if err != nil {
		return nil, err
	}

	u := ctx.Value("user").(models.User)
	if g.Owner.ID == u.ID {
		return nil, ErrGameUserExists
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
	if !g.UserInGame(u) {
		return nil, ErrGameUserNotExists
	}

	if g.CurrentPlayer.Player.ID != u.ID {
		return nil, ErrGameNotTurnUser
	}

	if err = g.MakeStep(c); err != nil {
		return nil, ErrGameCellOccupied
	}

	if err = s.UpdateGame(ctx, g); err != nil {
		return nil, err
	}

	return g, nil
}

func (s *Service) Subscribe(ctx context.Context, conn *websocket.Conn, gameID string) {
	gameUpdateCh := s.l.Join(gameID)
	for {
		select {
		case g := <-gameUpdateCh:
			b, err := json.Marshal(g)
			if err != nil {
				log.Println(err)
				continue
			}

			if err := conn.WriteMessage(websocket.TextMessage, b); err != nil {
				log.Println(err)
			}
		case <-ctx.Done():
			return
		}
	}
}

func (s *Service) HandleWSMessage(ctx context.Context, gameID string, rawCoord []byte) error {
	var coord models.Coord

	if err := json.Unmarshal(rawCoord, &coord); err != nil {
		log.Println("json unmarhsal hadle ws message:", err)
		return err
	}

	_, err := s.MakeStep(ctx, gameID, coord)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) getGameWithState(ctx context.Context, gameID string, state int16) (*models.Game, error) {
	g, err := s.GetGameByID(ctx, gameID)
	if err != nil {
		return nil, err
	}

	if g.CurrentState != state {
		return nil, ErrGameNotTurnUser
	}

	return g, nil
}
