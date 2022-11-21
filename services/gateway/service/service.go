package service

import (
	"errors"

	"github.com/Drozd0f/ttto-go/services/gateway/clients"
)

var (
	ErrGrpcBadError = errors.New("grpc bad error")
)

type Service struct {
	ac clients.AuthClient
}

func New(ac clients.AuthClient) *Service {
	return &Service{ac: ac}
}
