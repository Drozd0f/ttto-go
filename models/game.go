package models

import (
	"github.com/google/uuid"
)

type Field [3][3]string

const (
	StatePending = iota
	StateInGame
	StateDone
)

type Game struct {
	ID              uuid.UUID     `json:"id"`
	OwnerID         uuid.UUID     `json:"owner_id"`
	OpponentID      uuid.NullUUID `json:"opponent_id"`
	CurrentPlayerID uuid.NullUUID `json:"current_player_id"`
	StepCount       int32         `json:"step_count"`
	WinnerID        uuid.NullUUID `json:"winner_id"`
	Field           Field         `json:"field"`
	CurrentState    int16         `json:"current_state"`
}
