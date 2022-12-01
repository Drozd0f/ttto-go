package main

import (
	"fmt"

	"github.com/Drozd0f/ttto-go/services/gateway/service"
	"github.com/urfave/cli/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/Drozd0f/ttto-go/services/gateway/clients"
	"github.com/Drozd0f/ttto-go/services/gateway/conf"
	"github.com/Drozd0f/ttto-go/services/gateway/server"
)

func runGateway(c *cli.Context) error {
	cfg, err := conf.New()
	if err != nil {
		return fmt.Errorf("conf new: %w", err)
	}

	conn, err := grpc.Dial(cfg.AuthAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return fmt.Errorf("grpc dial: %w", err)
	}
	defer conn.Close()

	ac := clients.NewAuthClient(conn)
	s := service.New(ac)

	gtw, err := server.New(cfg, s)
	if err != nil {
		return fmt.Errorf("server new: %w", err)
	}

	if err = gtw.Run(cfg.Addr); err != nil {
		return fmt.Errorf("service gateway run: %w", err)
	}

	return nil
}
