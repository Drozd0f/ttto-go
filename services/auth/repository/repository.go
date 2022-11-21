package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Drozd0f/ttto-go/services/auth/db"
)

const uniqueConstraintCode = "23505"

var (
	ErrUniqueConstraint = errors.New("unique constraint")
	ErrNoResult = errors.New("no result")
)

type Repository struct {
	conn *sql.DB
	q    *db.Queries
}

func New(ctx context.Context, dbURI string) (*Repository, error) {
	conn, err := sql.Open("postgres", dbURI)
	if err != nil {
		return nil, fmt.Errorf("sql open: %w", err)
	}

	q, err := db.Prepare(ctx, conn)
	if err != nil {
		return nil, fmt.Errorf("db prepare: %w", err)
	}

	return &Repository{
		conn: conn,
		q:    q,
	}, nil
}
