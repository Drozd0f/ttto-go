package models

import (
	"encoding/json"

	"github.com/google/uuid"
)

type Player struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Mark     string    `json:"mark"`
}

type NullPlayer struct {
	Player Player
	Valid  bool
}

func (np NullPlayer) GetID() uuid.NullUUID {
	return uuid.NullUUID{
		UUID:  np.Player.ID,
		Valid: np.Valid,
	}
}

func (np NullPlayer) MarshalJSON() ([]byte, error) {
	if np.Valid {
		return json.Marshal(np.Player)
	}

	return []byte("null"), nil
}
