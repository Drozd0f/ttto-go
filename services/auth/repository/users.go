package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/Drozd0f/ttto-go/services/auth/schemes"
	"github.com/lib/pq"
)

func (r *Repository) CreateUser(ctx context.Context, user *schemes.User) error {
	err := r.q.CreateUser(ctx, user.UserToCreateUserParam())
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
