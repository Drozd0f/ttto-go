package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

func Populate(prefix string, spec any) error {
	if err := envconfig.Process(prefix, spec); err != nil {
		return fmt.Errorf("process config: %w", err)
	}

	return nil
}
