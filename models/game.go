package models

import (
	"database/sql"
	"encoding/json"
	"log"

	"github.com/Drozd0f/ttto-go/db"
	"github.com/google/uuid"
)

type Field [3][3]string

const (
	StatePending = iota
	StateInGame
	StateDone
)

type DBGame interface {
	db.CreateGameRow | db.GamesWithUsername
}

type GameWithUsernameSlice []GameWithUsername

type GameWithUsername struct {
	ID                uuid.UUID      `json:"id"`
	OwnerID           uuid.UUID      `json:"owner_id"`
	OpponentID        uuid.NullUUID  `json:"opponent_id"`
	CurrentPlayerID   uuid.NullUUID  `json:"current_player_id"`
	StepCount         int32          `json:"step_count"`
	WinnerID          uuid.NullUUID  `json:"winner_id"`
	Field             Field          `json:"field"`
	CurrentState      int16          `json:"current_state"`
	OwnerName         sql.NullString `json:"owner_name"`
	OpponentName      sql.NullString `json:"opponent_name"`
	CurrentPlayerName sql.NullString `json:"current_player_name"`
	WinnerName        sql.NullString `json:"winner_name"`
}

func GameWithUsernameSliceFromDB(storGames []db.GamesWithUsername) GameWithUsernameSlice {
	games := make(GameWithUsernameSlice, 0, len(storGames))
	for _, g := range storGames {
		games = append(games, *GameWithUsernameFromDB(g))
	}

	return games
}

func NewGameWithUsernameFromDB(g db.CreateGameRow) *GameWithUsername {
	var f Field
	if err := json.Unmarshal(g.Field.RawMessage, &f); err != nil {
		log.Println(err.Error())
	}

	return &GameWithUsername{
		ID:                g.ID,
		OwnerID:           g.OwnerID,
		OpponentID:        g.OpponentID,
		CurrentPlayerID:   g.CurrentPlayerID,
		StepCount:         g.StepCount,
		WinnerID:          g.WinnerID,
		Field:             f,
		CurrentState:      g.CurrentState,
		OwnerName:         g.OwnerName,
		OpponentName:      g.OpponentName,
		CurrentPlayerName: g.CurrentPlayerName,
		WinnerName:        g.WinnerName,
	}
}

func GameWithUsernameFromDB(g db.GamesWithUsername) *GameWithUsername {
	var f Field
	if err := json.Unmarshal(g.Field.RawMessage, &f); err != nil {
		log.Println(err.Error())
	}

	return &GameWithUsername{
		ID:                g.ID,
		OwnerID:           g.OwnerID,
		OpponentID:        g.OpponentID,
		CurrentPlayerID:   g.CurrentPlayerID,
		StepCount:         g.StepCount,
		WinnerID:          g.WinnerID,
		Field:             f,
		CurrentState:      g.CurrentState,
		OwnerName:         g.OwnerName,
		OpponentName:      g.OpponentName,
		CurrentPlayerName: g.CurrentPlayerName,
		WinnerName:        g.WinnerName,
	}
}
