package commands

import (
	"embed"
	"fmt"

	"github.com/Drozd0f/ttto-go/pkg/config"
	"github.com/Drozd0f/ttto-go/pkg/logger"
	"github.com/Drozd0f/ttto-go/pkg/migrator"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/urfave/cli/v2"
)

func Migrate(appName string, migrations embed.FS) *cli.Command {
	return &cli.Command{
		Name:   "migrate",
		Usage:  "Run database migrations",
		Action: actionMigrate(appName, migrations),
	}
}

func actionMigrate(appName string, migrations embed.FS) func(*cli.Context) error {
	return func(*cli.Context) error {
		c := config.NewMigrateConfig()
		if err := config.Populate(appName, c); err != nil {
			return fmt.Errorf("%s failed populate migration config: %w", appName, err)
		}

		logg, err := logger.NewLogger(c.Env)
		if err != nil {
			return fmt.Errorf("%s failed logger create: %w", appName, err)
		}

		if err := migrator.Migrate(migrations, c.DbConfig.ConnectionUrl(), logg); err != nil {
			return fmt.Errorf("%s failed migrate: %w", appName, err)
		}

		return nil
	}
}
