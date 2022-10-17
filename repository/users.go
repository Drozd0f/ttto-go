package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"

	"github.com/Drozd0f/ttto-go/db"
)

var ErrUserNotFound = errors.New("user not found")

func (r *Repository) GetUserByID(ctx context.Context, id uuid.UUID) (db.User, error) {
	u, err := r.q.GetUserByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return db.User{}, ErrUserNotFound
		}

		return db.User{}, fmt.Errorf("queries get user by id: %w", err)
	}

	return u, nil
}
