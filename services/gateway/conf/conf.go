package conf

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Addr   string
	Debug  bool
	AuthAddr string `split_words:"true"`
}

func New() (*Config, error) {
	var c Config

	if err := envconfig.Process("gateway", &c); err != nil {
		return nil, fmt.Errorf("envconfig process: %w", err)
	}

	fmt.Println(c)

	return &c, nil
}
