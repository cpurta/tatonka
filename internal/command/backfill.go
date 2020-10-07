package command

import (
	"github.com/cpurta/tatanka/internal/runner"
	"github.com/urfave/cli/v2"
)

// BackfillCommand allows for historical data to be manually
// back filled into a configured Cassandra cluster
// for trading simulations
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
