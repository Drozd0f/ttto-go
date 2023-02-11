package migrator

import (
	"embed"
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"go.uber.org/zap"
)

func Migrate(migrations embed.FS, dbURI string, logg *zap.Logger) error {
	src, err := iofs.New(migrations, ".")
	if err != nil {
		return fmt.Errorf("iofs new: %w", err)
	}

	m, err := migrate.NewWithSourceInstance("iofs", src, dbURI)
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

	logg.Info("migration applied",
		zap.Uint("current version", v),
		zap.Bool("is dirty", d),
	)

	return nil
}
