package config

import (
	"testing"

	"github.com/Drozd0f/csv-app/test/containers"
	"github.com/Drozd0f/ttto-go/pkg/config"
	"github.com/Drozd0f/ttto-go/pkg/env"
)

func NewTestConfig(t *testing.T, td *containers.TestDatabase) *Config {
	return &Config{
		Env:          env.EnvTest,
		Secret:       "secret",
		ServerConfig: ServerConfig{},
		DbConfig: config.DbConfig{
			User:     "postgres",
			Password: "postgres",
			Name:     "postgres",
			Host:     "localhost",
			Port:     "5432",
		},
	}
}
