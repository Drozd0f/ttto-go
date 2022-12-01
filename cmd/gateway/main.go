package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Action: runGateway,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
