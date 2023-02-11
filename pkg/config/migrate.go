package config

import "github.com/Drozd0f/ttto-go/pkg/env"

type MigrateConfig struct {
	Env      env.EnvType `default:"prod" required:"false"`
	DbConfig DbConfig
}

func NewMigrateConfig() *MigrateConfig {
	return &MigrateConfig{}
}
