package server

import (
	"context"
	"errors"

	"github.com/Drozd0f/ttto-go/gen/proto/auth"
	"github.com/Drozd0f/ttto-go/services/auth/schemes"
	"github.com/Drozd0f/ttto-go/services/auth/service"
	validation "github.com/go-ozzo/ozzo-validation"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) CreateUser(ctx context.Context, cur *auth.CreateUserRequest) (*auth.CreateUserResponse, error) {
	if err := s.service.CreateUser(ctx, schemes.UserFromCreateUserRequest(cur)); err != nil {
		var validErr validation.Errors
		switch {
		case errors.As(err, &validErr):
			return nil, status.Error(codes.InvalidArgument, validErr.Error())
		case errors.Is(err, service.ErrUserAlreadyExists):
			return nil, status.Error(codes.AlreadyExists, err.Error())
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}

	}

	return &auth.CreateUserResponse{}, nil
}

func (s *Server) LoginUser(ctx context.Context, lur *auth.LoginUserRequest) (*auth.LoginUserResponse, error) {
	token, err := s.service.LoginUser(ctx, schemes.UserFromLoginUserRequest(lur))
	if err != nil {
		var validErr validation.Errors
		switch {
		case errors.As(err, &validErr):
			return nil, status.Error(codes.InvalidArgument, validErr.Error())
		case errors.Is(err, service.ErrUserNotExists):
			return nil, status.Error(codes.NotFound, err.Error())
		case errors.Is(err, service.ErrUserInvalidData):
			return nil, status.Error(codes.InvalidArgument, err.Error())
		default:
			return nil, status.Error(codes.Internal, err.Error())
		}

	}

	return &auth.LoginUserResponse{Token: token}, nil
}
