package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Drozd0f/ttto-go/db"
	"github.com/Drozd0f/ttto-go/models"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

func (r *Repository) CreateUser(ctx context.Context, u *models.User) error {
	err := r.q.CreateUser(ctx, db.CreateUserParams{
		ID:       uuid.New(),
		Username: u.Name,
		Password: u.Password,
	})
	if err != nil {
		var pgError *pq.Error
		if errors.As(err, &pgError) {
			if pgError.Code == uniqueConstraintCode {
				return fmt.Errorf("user exist: %w", ErrUniqueConstraint)
			}
		}

		return err
	}

	return nil
}

func (r *Repository) GetUserByName(ctx context.Context, u *models.User) (db.User, error) {
	dbUser, err := r.q.GetUserByUsername(ctx, u.Name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return db.User{}, ErrNoResult
		}

		return db.User{}, err
	}

	return dbUser, nil
}
