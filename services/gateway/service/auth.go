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
	ErrUserNotExists     = errors.New("user not exist")
)

func (s *Service) CreateUser(ctx context.Context, cur *auth.CreateUserRequest) error {
	_, err := s.ac.CreateUser(ctx, cur)
	if err != nil {
		stat, ok := status.FromError(err)
		if ok {
			switch {
			case stat.Code() == codes.InvalidArgument:
				return ErrValidation
			case stat.Code() == codes.AlreadyExists:
				return ErrUserAlreadyExists
			}
			return fmt.Errorf("auth service create user: %s", stat.Message())
		}

		return fmt.Errorf("auth service status from error: %w", ErrGrpcBadError)
	}

	return nil
}

func (s *Service) LoginUser(ctx context.Context, lur *auth.LoginUserRequest) (string, error) {
	logUser, err := s.ac.LoginUser(ctx, lur)
	if err != nil {
		stat, ok := status.FromError(err)
		if ok {
			fmt.Println(stat.Code())
			switch {
			case stat.Code() == codes.InvalidArgument:  //TODO: validation error and invalid data in this case refactor to error with details
				return "", ErrValidation
			case stat.Code() == codes.NotFound:
				return "", ErrUserNotExists
			}
			return "", fmt.Errorf("auth service login user: %s", stat.Message())
		}

		return "", fmt.Errorf("auth service status from error: %w", ErrGrpcBadError)
	}


	return logUser.Token, err
}
