// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package db

import (
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID
	Username string
	Password string
}