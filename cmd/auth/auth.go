package main

import (
	"fmt"
	"net"

	"github.com/Drozd0f/ttto-go/gen/proto/auth"
	"github.com/urfave/cli/v2"
	"google.golang.org/grpc"

	"github.com/Drozd0f/ttto-go/services/auth/conf"
	"github.com/Drozd0f/ttto-go/services/auth/repository"
	"github.com/Drozd0f/ttto-go/services/auth/server"
	"github.com/Drozd0f/ttto-go/services/auth/service"
)

func runAuth(c *cli.Context) error {
	cfg, err := conf.New()
	if err != nil {
		return fmt.Errorf("conf new: %w", err)
	}

	rep, err := repository.New(c.Context, cfg.DBURI)
	if err != nil {
		return fmt.Errorf("repository new: %w", err)
	}

	s := server.New(cfg, service.New(rep, cfg))
	g := grpc.NewServer()
	auth.RegisterAuthServer(g, s)

	lis, err := net.Listen("tcp", cfg.Addr)
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}

	if err = g.Serve(lis); err != nil {
		return fmt.Errorf("grpc server serve: %w", err)
	}

	return nil
}
