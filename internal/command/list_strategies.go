package command

import (
	"github.com/cpurta/tatanka/internal/runner"
	"github.com/urfave/cli/v2"
)

func ListStrategiesCommand() *cli.Command {
	var (
		listStrategiesRunner = &runner.ListStrategiesRunner{}
	)

	cmd := &cli.Command{
		Name:   "list-strategies",
		Usage:  "list available strategies",
		Flags:  []cli.Flag{},
		Action: listStrategiesRunner.Run,
	}

	return cmd
}
