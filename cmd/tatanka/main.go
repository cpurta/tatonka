package main

import (
	"fmt"
	"os"

	"github.com/cpurta/tatanka/internal/command"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "tatanka",
		Usage: "command-line cryptocurrency trading bot written in Golang",
		Commands: []*cli.Command{
			command.BackfillCommand(),
			command.BalanceCommand(),
			command.ListSelectorCommand(),
			command.ListStrategiesCommand(),
			command.SimCommand(),
			command.TradeCommand(),
		},
		Version: "v0.0.1",
		Authors: []*cli.Author{
			{
				Name:  "Chris Purta",
				Email: "cpurta@gmail.com",
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println("error running program:", err.Error())
		os.Exit(1)
	}
}
