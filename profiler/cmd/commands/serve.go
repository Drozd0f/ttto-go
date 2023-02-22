package commands

import (
	"fmt"

	"github.com/Drozd0f/ttto-go/pkg/logger"
	"github.com/Drozd0f/ttto-go/profiler/config"
	"github.com/Drozd0f/ttto-go/profiler/repository"
	"github.com/Drozd0f/ttto-go/profiler/server"
	"github.com/Drozd0f/ttto-go/profiler/server/servers"
	"github.com/Drozd0f/ttto-go/profiler/service"
	"github.com/urfave/cli/v2"
	"golang.org/x/sync/errgroup"
)

func Serve(appName string) *cli.Command {
	return &cli.Command{
		Name:   "serve",
		Usage:  "Run profiler server",
		Action: actionServe(appName),
	}
}

func actionServe(appName string) func(c *cli.Context) error {
	return func(c *cli.Context) error {
		conf := config.New()
		if err := config.Populate(conf); err != nil {
			return fmt.Errorf("%s failed populate server config: %w", appName, err)
		}

		logg, err := logger.NewLogger(conf.Env)
		if err != nil {
			return fmt.Errorf("%s failed create logger: %w", appName, err)
		}
		// defer restore()

		repo, err := repository.NewRepository(c.Context, conf.DbConfig.ConnectionUrl())
		if err != nil {
			return fmt.Errorf("%s failed create repository: %w", appName, err)
		}

		s := service.New(repo, conf)

		serv := server.New(logg, conf, s)
		eg, ctx := errgroup.WithContext(c.Context)
		eg.Go(func() error { return servers.NewGrpcServer(ctx, logg, conf, serv) })
		eg.Go(func() error { return servers.NewHttpServer(ctx, logg, conf) })

		if err := eg.Wait(); err != nil {
			return fmt.Errorf("%s wait server: %w", appName, err)
		}

		return nil
	}
}
