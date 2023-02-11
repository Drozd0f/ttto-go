package main

import (
	"log"
	"os"

	"github.com/Drozd0f/ttto-go/profiler/cmd/commands"
	"github.com/Drozd0f/ttto-go/profiler/config"
	"github.com/Drozd0f/ttto-go/profiler/repository/migrations"
	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = config.AppName
	app.Usage = "Profiler service"
	app.Commands = cli.Commands{
		commands.Migrate(config.AppName, migrations.Migrations),
		commands.Serve(config.AppName),
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
