// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package db

import (
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Password string    `json:"password"`
}
