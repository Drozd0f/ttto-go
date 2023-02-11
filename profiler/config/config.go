package config

import (
	"net"

	"github.com/Drozd0f/ttto-go/pkg/config"
	"github.com/Drozd0f/ttto-go/pkg/env"
)

const AppName = "profiler"

type Config struct {
	Env    env.EnvType `default:"prod" required:"false"`
	Secret string      `default:"secret" required:"true"`

	ServerConfig
	DbConfig config.DbConfig
}

type ServerConfig struct {
	Host     string `default:"localhost"`
	HttpPort string `default:"8080" split_words:"true"`
	GrpcPort string `default:"20000" split_words:"true"`
}

func New() *Config {
	return &Config{}
}

func (c Config) HttpAdress() string {
	return net.JoinHostPort(c.Host, c.HttpPort)
}
func (c Config) GrpcAdress() string {
	return net.JoinHostPort(c.Host, c.GrpcPort)
}

func Populate(c *Config) error {
	return config.Populate(AppName, c)
}
