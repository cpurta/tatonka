package command

import (
	"github.com/cpurta/tatanka/internal/runner"
	"github.com/urfave/cli/v2"
)

func SimCommand() *cli.Command {
	var (
		simRunner = &runner.SimRunner{}
	)

	cmd := &cli.Command{
		Name:  "sim",
		Usage: "run a simulation on backfilled data",
		Flags: []cli.Flag{
			&cli.PathFlag{
				Name:        "config",
				Usage:       "path to optional config overrides file",
				Destination: &simRunner.ConfigFile,
				Value:       "/etc/tatanka/config.yaml",
			},
			&cli.StringFlag{
				Name:        "strategy",
				Usage:       "strategy to use",
				Destination: &simRunner.Strategy,
				Value:       "trend_ema",
			},
			&cli.StringFlag{
				Name:        "order_type",
				Usage:       "order type to use (maker/taker)",
				Destination: &simRunner.OrderType,
				Value:       "maker",
			},
			&cli.PathFlag{
				Name:        "filename",
				Usage:       "filename for the result output (ex: result.html). \"none\" to disable",
				Destination: &simRunner.Filename,
			},
			&cli.Int64Flag{
				Name:        "days",
				Usage:       "set duration by day count",
				Destination: &simRunner.Days,
				Value:       14,
			},
			&cli.Int64Flag{
				Name:        "currency_capital",
				Usage:       "amount of start capital in currency",
				Destination: &simRunner.CurrencyCapital,
				Value:       1000,
			},
			&cli.Int64Flag{
				Name:        "asset_capital",
				Usage:       "amount of start capital in asset",
				Destination: &simRunner.AssetCapital,
				Value:       0,
			},
			&cli.Float64Flag{
				Name:        "avg_slippage_pct",
				Usage:       "avg. amount of slippage to apply to trades",
				Destination: &simRunner.AvgSlippagePct,
				Value:       0.045,
			},
			&cli.Int64Flag{
				Name:        "buy_pct",
				Usage:       "buy with this % of currency balance",
				Destination: &simRunner.BuyPct,
				Value:       99,
			},
			&cli.Int64Flag{
				Name:        "sell_pct",
				Usage:       "sell with this % of asset balance",
				Destination: &simRunner.SellPct,
				Value:       99,
			},
			&cli.Int64Flag{
				Name:        "markdown_buy_pct",
				Usage:       "% to mark down buy price",
				Destination: &simRunner.MarkdownBuyPct,
				Value:       0,
			},
			&cli.Int64Flag{
				Name:        "markdown_sell_pct",
				Usage:       "% to mark up sell price",
				Destination: &simRunner.MarkupSellPct,
				Value:       0,
			},
			&cli.Int64Flag{
				Name:        "order_adjust_time",
				Usage:       "adjust bid/ask on this interval to keep orders competitive in ms",
				Destination: &simRunner.OrderAdjustTime,
				Value:       5000,
			},
			&cli.Int64Flag{
				Name:        "order_poll_time",
				Usage:       "poll order status on this interval in ms",
				Destination: &simRunner.OrderPollTime,
				Value:       5000,
			},
			&cli.Float64Flag{
				Name:        "sell_cancel_pct",
				Usage:       "cancels the sale if the price is between this percentage (for more or less)",
				Destination: &simRunner.SellCancelPct,
			},
			&cli.Float64Flag{
				Name:        "sell_stop_pct",
				Usage:       "sell if price drops below this % of bought price",
				Destination: &simRunner.SellStopPct,
				Value:       0.0,
			},
			&cli.Float64Flag{
				Name:        "buy_stop_pct",
				Usage:       "buy if price surges above this % of sold price",
				Destination: &simRunner.BuyStopPct,
				Value:       0.0,
			},
			&cli.Int64Flag{
				Name:        "profit_stop_enable_pct",
				Usage:       "enable trailing sell stop when reaching this % profit",
				Destination: &simRunner.ProfitStopEnablePct,
				Value:       0,
			},
			&cli.Int64Flag{
				Name:        "profit_stop_pct",
				Usage:       "maintain a trailing stop this % below the high-water mark of profit",
				Destination: &simRunner.ProfitStopEnablePct,
				Value:       1,
			},
			&cli.Int64Flag{
				Name:        "max_sell_loss_pct",
				Usage:       "avoid selling at a loss pct under this float",
				Destination: &simRunner.MaxSellLossPct,
				Value:       99,
			},
			&cli.Int64Flag{
				Name:        "max_buy_loss_pct",
				Usage:       "avoid buying at a loss pct over this float",
				Destination: &simRunner.MaxBuyLossPct,
				Value:       99,
			},
			&cli.Int64Flag{
				Name:        "max_slippage_pct",
				Usage:       "avoid selling at a slippage pct above this float",
				Destination: &simRunner.MaxSlippagePct,
				Value:       5,
			},
			&cli.Int64Flag{
				Name:        "rsi_periods",
				Usage:       "number of periods to calculate RSI at",
				Destination: &simRunner.RSIPeriods,
				Value:       14,
			},
			&cli.BoolFlag{
				Name:        "exact_buy_orders",
				Usage:       "instead of only adjusting maker buy when the price goes up, adjust it if price has changed at all",
				Destination: &simRunner.ExactBuyOrders,
			},
			&cli.BoolFlag{
				Name:        "exact_sell_orders",
				Usage:       "instead of only adjusting maker sell when the price goes down, adjust it if price has changed at all",
				Destination: &simRunner.ExactSellOrders,
			},
			&cli.BoolFlag{
				Name:        "disable_options",
				Usage:       "disable printing of options",
				Destination: &simRunner.DisableOptions,
			},
			&cli.Int64Flag{
				Name:        "quarentine_time",
				Usage:       "For loss trade, set quarentine time for cancel buys in minutes",
				Destination: &simRunner.QuarentineTime,
				Value:       10,
			},
			&cli.BoolFlag{
				Name:        "enable_stats",
				Usage:       "enable printing order stats",
				Destination: &simRunner.DisableOptions,
			},
			&cli.Int64Flag{
				Name:        "backtester_generation",
				Usage:       "creates a json file in simulations with the generation number",
				Destination: &simRunner.BacktesterGeneration,
				Value:       -1,
			},
			&cli.BoolFlag{
				Name:        "verbose",
				Usage:       "print status lines on every period",
				Destination: &simRunner.Verbose,
			},
			&cli.BoolFlag{
				Name:        "silent",
				Usage:       "only output on completion (can speed up sim)",
				Destination: &simRunner.Silent,
			},
		},
		Action: simRunner.Run,
	}

	return cmd
}
