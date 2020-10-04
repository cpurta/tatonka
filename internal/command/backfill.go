package command

import (
	"github.com/cpurta/tatanka/internal/runner"
	"github.com/urfave/cli/v2"
)

// BackfillCommand returns the history of trades
// with the option to different path for config file,
// number of days with occurred or with debug
// CLI Command
func BackfillCommand() *cli.Command {
	var (
		backfillRunner = &runner.BackfillRunner{}
	)

	cmd := &cli.Command{
		Name:  "backfill",
		Usage: "download historical trades for analysis",
		Flags: []cli.Flag{
			&cli.PathFlag{
				Name:        "config",
				Usage:       "path to optional config overrides file",
				Destination: &backfillRunner.ConfigFile,
				Value:       "/etc/tatanka/config.yaml",
			},
			&cli.Int64Flag{
				Name:        "days",
				Usage:       "number of days to acquire",
				Destination: &backfillRunner.Days,
				Value:       1,
			},
			&cli.BoolFlag{
				Name:        "debug",
				Usage:       "output detailed debug info",
				Destination: &backfillRunner.Debug,
			},
		},
		Action: backfillRunner.Run,
	}

	return cmd
}
