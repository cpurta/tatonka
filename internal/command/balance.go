package command

import (
	"github.com/cpurta/tatonka/internal/runner"
	"github.com/urfave/cli/v2"
)

// BalanceCommand returns the current account balances
// with the option to change to another currency
// CLI Command
func BalanceCommand() *cli.Command {
	var (
		balanceRunner = &runner.BalanceRunner{}
	)

	cmd := &cli.Command{
		Name:  "balance",
		Usage: "get asset and currency balance from the exchange",
		Flags: []cli.Flag{
			&cli.PathFlag{
				Name:        "config",
				Usage:       "path to optional config overrides file",
				Destination: &balanceRunner.ConfigFile,
				Value:       "/etc/tatonka/config.yaml",
			},
			&cli.BoolFlag{
				Name:        "calculate_currency",
				Usage:       "show the full balance in another currency",
				Destination: &balanceRunner.CalculateCurrency,
			},
		},
		Action: balanceRunner.Run,
	}

	return cmd
}
