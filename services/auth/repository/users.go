package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Drozd0f/ttto-go/services/auth/db"
	"github.com/Drozd0f/ttto-go/services/auth/schemes"
	"github.com/lib/pq"
)

func (r *Repository) CreateUser(ctx context.Context, u schemes.User) error {
	err := r.q.CreateUser(ctx, u.UserToCreateUserParam())
	if err != nil {
		var pgError *pq.Error
		if errors.As(err, &pgError) {
			if pgError.Code == uniqueConstraintCode {
				return fmt.Errorf("queries create user: %w", ErrUniqueConstraint)
			}
		}

		return err
	}

	return nil
}

func (r *Repository) GetUserByName(ctx context.Context, u schemes.User) (db.User, error) {
	dbUser, err := r.q.GetUserByUsername(ctx, u.Username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return db.User{}, ErrNoResult
		}

		return db.User{}, fmt.Errorf("queries get user by username: %w", err)
	}

	return dbUser, nil
}
