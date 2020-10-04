package command

import (
	"github.com/cpurta/tatanka/internal/runner"
	"github.com/urfave/cli/v2"
)

// ListSelectorCommand Print all available exchange options
// CLI Command
func ListSelectorCommand() *cli.Command {
	var (
		listSelectorRunner = &runner.ListSelectorRunner{}
	)

	cmd := &cli.Command{
		Name:   "list-selectors",
		Usage:  "list available selectors",
		Flags:  []cli.Flag{},
		Action: listSelectorRunner.Run,
	}

	return cmd
}
