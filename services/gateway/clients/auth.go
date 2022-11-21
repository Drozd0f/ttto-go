package clients

import (
	"github.com/Drozd0f/ttto-go/gen/proto/auth"
	"google.golang.org/grpc"
)

type AuthClient interface {
	auth.AuthClient
}

func NewAuthClient(cc grpc.ClientConnInterface) AuthClient {
	return auth.NewAuthClient(cc)
}
