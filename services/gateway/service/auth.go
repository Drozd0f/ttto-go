package service

import (
	"context"
	"errors"
	"fmt"

	"github.com/Drozd0f/ttto-go/gen/proto/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrValidation = errors.New("validation error")
	ErrUserAlreadyExists   = errors.New("user already exist")
)

func (s *Service) CreateUser(ctx context.Context, req *auth.CreateUserRequest) error {
	_, err := s.ac.CreateUser(ctx, req)

	stat, ok := status.FromError(err)
	if ok {
		switch {
		case stat.Code() == codes.InvalidArgument:
			return ErrValidation
		case stat.Code() == codes.AlreadyExists:
			return ErrUserAlreadyExists
		default:
			return fmt.Errorf("auth service create user: %s", stat.Message())
		}
	}

	return nil
}
