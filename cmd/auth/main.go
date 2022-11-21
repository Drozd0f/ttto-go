package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name: "auth",
		Commands: []*cli.Command{
			{
				Name:   "run",
				Action: runAuth,
			},
			{
				Name:   "migrate",
				Action: runMigrate,
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
