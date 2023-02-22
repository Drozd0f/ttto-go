package config

import (
	"net"
	"net/url"
)

type DbConfig struct {
	User           string `required:"true"`
	Password       string `required:"true"`
	Name           string `required:"true"`
	Host           string `required:"true"`
	Port           string `required:"true"`
	MaxOpenConn    int    `default:"5" required:"true"`
	ConMaxLifeTime string `default:"30m" required:"true"`
}

func (d DbConfig) ConnectionUrl() string {
	u := url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(d.User, d.Password),
		Host:   net.JoinHostPort(d.Host, d.Port),
		Path:   d.Name,
		RawQuery: url.Values{
			"sslmode": []string{"disable"},
		}.Encode(),
	}

	return u.String()
}
