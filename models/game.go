package models

import (
	"encoding/json"
	"errors"
	"log"
	"math"
	"math/rand"

	"github.com/Drozd0f/ttto-go/db"
	"github.com/google/uuid"
)

const (
	StatePending = iota
	StateInGame
	StateDone
)

const (
	OwnerMark    = "X"
	OpponentMark = "0"
)

var winCondition = [8][3]int{
	{0, 1, 2},
	{3, 4, 5},
	{6, 7, 8},
	{0, 3, 6},
	{1, 4, 7},
	{2, 5, 8},
	{0, 4, 8},
	{2, 4, 6},
}

var ErrCellOccupied = errors.New("cell is already occupied")

type Field [3][3]string

type GameSlice []*Game

type Game struct {
	ID            uuid.UUID  `json:"id"`
	Owner         Player     `json:"owner"`
	Opponent      NullPlayer `json:"opponent"`
	CurrentPlayer NullPlayer `json:"current_player"`
	StepCount     int32      `json:"step_count"`
	Winner        NullPlayer `json:"winner"`
	Field         Field      `json:"field"`
	CurrentState  int16      `json:"current_state"`
}

func NewGameSliceFromDB(storGames []db.GamesWithUsername) GameSlice {
	games := make(GameSlice, 0, len(storGames))
	for _, g := range storGames {
		games = append(games, NewGameFromDB(g))
	}

	return games
}

func NewGameFromDB(g any) *Game {
	var f Field
	switch g := g.(type) {
	case db.CreateGameRow:
		if err := json.Unmarshal(g.Field, &f); err != nil {
			log.Println(err.Error())
		}
		game := &Game{
			ID: g.ID,
			Owner: Player{
				ID:       g.OwnerID,
				Username: g.OwnerName.String, // TODO: null string to string
				Mark:     OwnerMark,
			},
			StepCount:    g.StepCount,
			Field:        f,
			CurrentState: g.CurrentState,
		}
		if g.OpponentID.Valid {
			game.Opponent = NullPlayer{
				Player: Player{
					ID:       g.OpponentID.UUID,
					Username: g.OpponentName.String,
					Mark:     OpponentMark,
				},
				Valid: true,
			}
		}
		if g.CurrentPlayerID.Valid {
			game.CurrentPlayer = NullPlayer{
				Player: Player{
					ID:       g.CurrentPlayerID.UUID,
					Username: g.CurrentPlayerName.String,
					Mark:     game.playerMark(g.CurrentPlayerID.UUID),
				},
				Valid: true,
			}
		}
		if g.WinnerID.Valid {
			game.Winner = NullPlayer{
				Player: Player{
					ID:       g.WinnerID.UUID,
					Username: g.WinnerName.String,
					Mark:     game.playerMark(g.WinnerID.UUID),
				},
				Valid: true,
			}
		}
		return game

	case db.GamesWithUsername:
		if err := json.Unmarshal(g.Field, &f); err != nil {
			log.Println(err.Error())
		}
		game := &Game{
			ID: g.ID,
			Owner: Player{
				ID:       g.OwnerID,
				Username: g.OwnerName.String, // TODO: null string to string
				Mark:     OwnerMark,
			},
			StepCount:    g.StepCount,
			Field:        f,
			CurrentState: g.CurrentState,
		}
		if g.OpponentID.Valid {
			game.Opponent = NullPlayer{
				Player: Player{
					ID:       g.OpponentID.UUID,
					Username: g.OpponentName.String,
					Mark:     OpponentMark,
				},
				Valid: true,
			}
		}
		if g.CurrentPlayerID.Valid {
			game.CurrentPlayer = NullPlayer{
				Player: Player{
					ID:       g.CurrentPlayerID.UUID,
					Username: g.CurrentPlayerName.String,
					Mark:     game.playerMark(g.CurrentPlayerID.UUID),
				},
				Valid: true,
			}
		}
		if g.WinnerID.Valid {
			game.Winner = NullPlayer{
				Player: Player{
					ID:       g.WinnerID.UUID,
					Username: g.WinnerName.String,
					Mark:     game.playerMark(g.WinnerID.UUID),
				},
				Valid: true,
			}
		}

		return game
	default:
		return &Game{}
	}
}

func (g *Game) SetOpponent(o User) {
	g.Opponent = NullPlayer{
		Player: Player{
			ID:       o.ID,
			Username: o.Username,
			Mark:     OpponentMark,
		},
		Valid: true,
	}
	g.CurrentState = StateInGame
	if rand.Intn(2) == 0 {
		g.CurrentPlayer = NullPlayer{
			Player: g.Owner,
			Valid:  true,
		}
		return
	}

	g.CurrentPlayer = g.Opponent
}

func (g *Game) Step(coords [2]int) error {
	row, col := coords[0], coords[1]
	if g.Field[row][col] != "" {
		return ErrCellOccupied
	}

	fieldSize := math.Pow(float64(len(g.Field)), 2.0)
	g.Field[row][col] = g.playerMark(g.CurrentPlayer.Player.ID)
	g.StepCount++
	if g.StepCount >= int32(math.Ceil(fieldSize/2)) && g.winCheck() {
		g.Winner = g.CurrentPlayer
		g.CurrentState = StateDone
		return nil
	}

	g.invertPlayer()
	if g.StepCount == int32(fieldSize) {
		g.CurrentState = StateDone
	}
	return nil
}

func (g *Game) playerMark(pID uuid.UUID) string {
	if g.Owner.ID == pID {
		return OwnerMark
	}
	return OpponentMark
}

func (g *Game) expandField() []string {
	fieldSize := len(g.Field)
	ef := make([]string, 0, fieldSize*fieldSize)
	for _, row := range g.Field {
		for _, col := range row {
			ef = append(ef, col)
		}
	}
	return ef
}

func (g *Game) winCheck() bool {
	ef := g.expandField()
	isWin := false
outer_loop:
	for _, cond := range winCondition {
		for _, cellIdx := range cond {
			if g.CurrentPlayer.Player.Mark != ef[cellIdx] {
				continue outer_loop
			}
		}
		isWin = true
		break
	}
	return isWin
}

func (g *Game) invertPlayer() {
	if g.CurrentPlayer.Player.ID == g.Owner.ID {
		g.CurrentPlayer.Player = g.Owner
		return
	}
	g.CurrentPlayer = g.Opponent
}

func (g *Game) MarshalField() []byte {
	b, err := json.Marshal(g.Field)
	if err != nil {
		log.Println(err.Error())
	}

	return b
}
