package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Drozd0f/ttto-go/schemes"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

const uniqueConstraintCode = "23505"

var (
	ErrUniqueConstraint = errors.New("unique constraint")
	ErrNoResult         = errors.New("no result")
)

type Repository struct {
	conn *sql.DB
	q    *Queries
}

func NewRepository(ctx context.Context, dbURI string) (*Repository, error) {
	conn, err := sql.Open("postgres", dbURI)
	if err != nil {
		return nil, fmt.Errorf("sql open: %w", err)
	}

	q, err := Prepare(ctx, conn)
	if err != nil {
		return nil, fmt.Errorf("db prepare: %w", err)
	}

	return &Repository{
		conn: conn,
		q:    q,
	}, nil
}

func (r *Repository) CreateUser(ctx context.Context, u *schemes.User) (*schemes.User, error) {
	dbUser, err := r.q.CreateUser(ctx, CreateUserParams{
		ID:       uuid.New(),
		Username: u.Username,
		Password: u.Password,
	})
	if err != nil {
		var pgError *pq.Error
		if errors.As(err, &pgError) {
			if pgError.Code == uniqueConstraintCode {
				return nil, fmt.Errorf("user exist: %w", ErrUniqueConstraint)
			}
		}

		return nil, err
	}

	return &schemes.User{
		ID:       dbUser.ID,
		Username: dbUser.Username,
	}, nil
}

func (r *Repository) GetUserByUsername(ctx context.Context, u *schemes.User) (*schemes.User, error) {
	dbUser, err := r.q.GetUserByUsername(ctx, u.Username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoResult
		}

		return nil, err
	}

	return &schemes.User{
		ID:       dbUser.ID,
		Username: dbUser.Username,
		Password: dbUser.Password,
	}, nil
}
