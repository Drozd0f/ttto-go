package service

import (
	"github.com/Drozd0f/ttto-go/services/gateway/clients"
)

type Service struct {
	ac clients.AuthClient
}

func New(ac clients.AuthClient) *Service {
	return &Service{ac: ac}
}
