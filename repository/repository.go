package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Drozd0f/ttto-go/db"
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
