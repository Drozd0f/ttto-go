package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name: "ttto",
		Commands: []*cli.Command{
			{
				Name:   "run-monolith",
				Action: runMonolithServer,
			},
			{
				Name:   "migrate-monolith",
				Action: runMonolithMigrate,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
