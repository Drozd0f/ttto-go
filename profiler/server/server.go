package server

import (
	"context"
	"errors"
	"net/http"

	profilerv1 "github.com/Drozd0f/ttto-go/gen/proto/go/profiler/v1"
	"github.com/Drozd0f/ttto-go/profiler/config"
	"github.com/Drozd0f/ttto-go/profiler/service"
	"github.com/Drozd0f/ttto-go/schemes"
	validation "github.com/go-ozzo/ozzo-validation"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	profilerv1.UnimplementedProfilerServiceServer

	logger  *zap.Logger
	c       *config.Config
	service *service.Service
}

var _ profilerv1.ProfilerServiceServer = (*Server)(nil)

func New(logger *zap.Logger, c *config.Config, service *service.Service) *Server {
	return &Server{
		logger:  logger,
		c:       c,
		service: service,
	}
}

func (s *Server) CreateUser(ctx context.Context, cur *profilerv1.CreateUserRequest) (*profilerv1.CreateUserResponse, error) {
	u := &schemes.User{
		Username: cur.GetUsername(),
		Password: cur.GetPassword(),
	}

	us, err := s.service.CreateUser(ctx, u)
	if err != nil {
		var validErr validation.Errors
		switch {
		case errors.As(err, &validErr):
			return nil, status.Error(codes.InvalidArgument, validErr.Error())
		case errors.Is(err, service.ErrUserAlreadyExists):
			return nil, status.Error(codes.AlreadyExists, err.Error())
		}

		s.logger.Error("service create user",
			zap.String("username", u.Username),
			zap.String("password", u.Password),
			zap.Error(err),
		)
		return nil, status.Error(codes.Internal, http.StatusText(http.StatusInternalServerError))
	}

	return &profilerv1.CreateUserResponse{
		Id:       us.ID.String(),
		Username: us.Username,
	}, nil
}
func (s *Server) GetUserByUsername(ctx context.Context, gur *profilerv1.GetUserByUsernameRequest) (*profilerv1.GetUserByUsernameResponse, error) {
	u := &schemes.User{
		Username: gur.GetUsername(),
	}

	us, err := s.service.GetUserByUsername(ctx, u)
	if err != nil {
		s.logger.Error("service get user",
			zap.String("username", u.Username),
			zap.Error(err),
		)
		return nil, status.Error(codes.Internal, http.StatusText(http.StatusInternalServerError))
	}

	return &profilerv1.GetUserByUsernameResponse{
		Id:       us.ID.String(),
		Username: us.Username,
		Password: us.Password,
	}, nil
}
