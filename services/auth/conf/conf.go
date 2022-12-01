package conf

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Addr   string
	DBURI  string
	Secret string
}

func New() (*Config, error) {
	var c Config

	if err := envconfig.Process("auth", &c); err != nil {
		return nil, fmt.Errorf("envconfig process: %w", err)
	}

	return &c, nil
}
