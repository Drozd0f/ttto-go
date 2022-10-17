package main

import (
	"fmt"

	"github.com/urfave/cli/v2"

	"github.com/Drozd0f/ttto-go/repository"
	"github.com/Drozd0f/ttto-go/server"
	"github.com/Drozd0f/ttto-go/service"
)

func runServer(c *cli.Context) error {
	r, err := repository.New(
		c.Context,
		"postgres://test:test@localhost:5432/test?sslmode=disable",
	)
	if err != nil {
		return fmt.Errorf("repository new: %w", err)
	}

	s := service.New(r)
	serv := server.New(s)

	if err = serv.Run(); err != nil {
		return fmt.Errorf("server run: %w", err)
	}

	return nil
}
