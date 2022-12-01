package main

import (
	"fmt"

	"github.com/Drozd0f/ttto-go/pkg/migrator"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/urfave/cli/v2"

	"github.com/Drozd0f/ttto-go/monolith/conf"
	"github.com/Drozd0f/ttto-go/monolith/db/migrations"
)

func runMonolithMigrate(*cli.Context) error {
	cfg, err := conf.New()
	if err != nil {
		return fmt.Errorf("conf new: %w", err)
	}

	if err = migrator.Migrate(migrations.Migrations, cfg.DBURI); err != nil {
		return fmt.Errorf("run migrate: %w", err)
	}

	return nil
}
