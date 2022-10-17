package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/urfave/cli/v2"

	"github.com/Drozd0f/ttto-go/db/migrations"
)

func runMigrate(*cli.Context) error {
	src, err := iofs.New(migrations.Migrations, ".")
	if err != nil {
		return fmt.Errorf("iofs new: %w", err)
	}

	m, err := migrate.NewWithSourceInstance("iofs", src, "postgres://test:test@localhost:5432/test?sslmode=disable")
	if err != nil {
		return fmt.Errorf("migrate new with source instance: %w", err)
	}

	if err = m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("migration up: %w", err)
	}

	v, d, err := m.Version()
	if err != nil {
		return fmt.Errorf("migration get version: %w", err)
	}

	log.Printf("current migration %d, is %t\n", v, d)

	return nil
}
