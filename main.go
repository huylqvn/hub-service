package main

import (
	"fmt"
	"hub-service/cmd/server"
	"hub-service/cmd/tool/migration"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "hub-service",
		Usage: "Hub CLI tool",
		Commands: []*cli.Command{
			{
				Name:  "server",
				Usage: "start server",
				Action: func(c *cli.Context) error {
					return server.Init()
				},
			},
			{
				Name:  "migration",
				Usage: "execute migration",
				Action: func(c *cli.Context) error {
					return migration.Run()
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
