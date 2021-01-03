package command

import (
	"github.com/cpurta/tatonka/internal/runner"
	"github.com/urfave/cli/v2"
)

func TradeCommand() *cli.Command {
	var (
		tradeRunner = &runner.TradeRunner{}
	)

	cmd := &cli.Command{
		Name:  "trade",
		Usage: "run trading bot against live market data",
		Flags: []cli.Flag{
			&cli.PathFlag{
				Name:        "config",
				Usage:       "path to optional config overrides file",
				Destination: &tradeRunner.ConfigFile,
				Value:       "/etc/tatonka/config.yaml",
			},
			&cli.StringFlag{
				Name:        "strategy",
				Usage:       "strategy to use",
				Destination: &tradeRunner.Strategy,
				Value:       "trend_ema",
			},
			&cli.StringFlag{
				Name:        "order_type",
				Usage:       "order type to use (maker/taker)",
				Destination: &tradeRunner.OrderType,
				Value:       "maker",
			},
			&cli.PathFlag{
				Name:        "filename",
				Usage:       "filename for the result output (ex: result.html). \"none\" to disable",
				Destination: &tradeRunner.Filename,
			},
			&cli.Int64Flag{
				Name:        "days",
				Usage:       "set duration by day count",
				Destination: &tradeRunner.Days,
				Value:       14,
			},
			&cli.Int64Flag{
				Name:        "currency_capital",
				Usage:       "amount of start capital in currency",
				Destination: &tradeRunner.CurrencyCapital,
				Value:       1000,
			},
			&cli.Int64Flag{
				Name:        "asset_capital",
				Usage:       "amount of start capital in asset",
				Destination: &tradeRunner.AssetCapital,
				Value:       0,
			},
			&cli.Float64Flag{
				Name:        "avg_slippage_pct",
				Usage:       "avg. amount of slippage to apply to trades",
				Destination: &tradeRunner.AvgSlippagePct,
				Value:       0.045,
			},
			&cli.Int64Flag{
				Name:        "buy_pct",
				Usage:       "buy with this % of currency balance",
				Destination: &tradeRunner.BuyPct,
				Value:       99,
			},
			&cli.Int64Flag{
				Name:        "sell_pct",
				Usage:       "sell with this % of asset balance",
				Destination: &tradeRunner.SellPct,
				Value:       99,
			},
			&cli.Int64Flag{
				Name:        "markdown_buy_pct",
				Usage:       "% to mark down buy price",
				Destination: &tradeRunner.MarkdownBuyPct,
				Value:       0,
			},
			&cli.Int64Flag{
				Name:        "markdown_sell_pct",
				Usage:       "% to mark up sell price",
				Destination: &tradeRunner.MarkupSellPct,
				Value:       0,
			},
			&cli.Int64Flag{
				Name:        "order_adjust_time",
				Usage:       "adjust bid/ask on this interval to keep orders competitive in ms",
				Destination: &tradeRunner.OrderAdjustTime,
				Value:       5000,
			},
			&cli.Int64Flag{
				Name:        "order_poll_time",
				Usage:       "poll order status on this interval in ms",
				Destination: &tradeRunner.OrderPollTime,
				Value:       5000,
			},
			&cli.Float64Flag{
				Name:        "sell_cancel_pct",
				Usage:       "cancels the sale if the price is between this percentage (for more or less)",
				Destination: &tradeRunner.SellCancelPct,
			},
			&cli.Float64Flag{
				Name:        "sell_stop_pct",
				Usage:       "sell if price drops below this % of bought price",
				Destination: &tradeRunner.SellStopPct,
				Value:       0.0,
			},
			&cli.Float64Flag{
				Name:        "buy_stop_pct",
				Usage:       "buy if price surges above this % of sold price",
				Destination: &tradeRunner.BuyStopPct,
				Value:       0.0,
			},
			&cli.Int64Flag{
				Name:        "profit_stop_enable_pct",
				Usage:       "enable trailing sell stop when reaching this % profit",
				Destination: &tradeRunner.ProfitStopEnablePct,
				Value:       0,
			},
			&cli.Int64Flag{
				Name:        "profit_stop_pct",
				Usage:       "maintain a trailing stop this % below the high-water mark of profit",
				Destination: &tradeRunner.ProfitStopEnablePct,
				Value:       1,
			},
			&cli.Int64Flag{
				Name:        "max_sell_loss_pct",
				Usage:       "avoid selling at a loss pct under this float",
				Destination: &tradeRunner.MaxSellLossPct,
				Value:       99,
			},
			&cli.Int64Flag{
				Name:        "max_buy_loss_pct",
				Usage:       "avoid buying at a loss pct over this float",
				Destination: &tradeRunner.MaxBuyLossPct,
				Value:       99,
			},
			&cli.Int64Flag{
				Name:        "max_slippage_pct",
				Usage:       "avoid selling at a slippage pct above this float",
				Destination: &tradeRunner.MaxSlippagePct,
				Value:       5,
			},
			&cli.Int64Flag{
				Name:        "rsi_periods",
				Usage:       "number of periods to calculate RSI at",
				Destination: &tradeRunner.RSIPeriods,
				Value:       14,
			},
			&cli.BoolFlag{
				Name:        "exact_buy_orders",
				Usage:       "instead of only adjusting maker buy when the price goes up, adjust it if price has changed at all",
				Destination: &tradeRunner.ExactBuyOrders,
			},
			&cli.BoolFlag{
				Name:        "exact_sell_orders",
				Usage:       "instead of only adjusting maker sell when the price goes down, adjust it if price has changed at all",
				Destination: &tradeRunner.ExactSellOrders,
			},
			&cli.BoolFlag{
				Name:        "disable_options",
				Usage:       "disable printing of options",
				Destination: &tradeRunner.DisableOptions,
			},
			&cli.Int64Flag{
				Name:        "quarentine_time",
				Usage:       "For loss trade, set quarentine time for cancel buys in minutes",
				Destination: &tradeRunner.QuarentineTime,
				Value:       10,
			},
			&cli.BoolFlag{
				Name:        "enable_stats",
				Usage:       "enable printing order stats",
				Destination: &tradeRunner.DisableOptions,
			},
			&cli.Int64Flag{
				Name:        "backtester_generation",
				Usage:       "creates a json file in simulations with the generation number",
				Destination: &tradeRunner.BacktesterGeneration,
				Value:       -1,
			},
			&cli.BoolFlag{
				Name:        "verbose",
				Usage:       "print status lines on every period",
				Destination: &tradeRunner.Verbose,
			},
			&cli.BoolFlag{
				Name:        "paper",
				Usage:       "use paper trading mode (no real trades will take place)",
				Destination: &tradeRunner.PaperTrade,
			},
			&cli.BoolFlag{
				Name:        "silent",
				Usage:       "only output on completion (can speed up sim)",
				Destination: &tradeRunner.Silent,
			},
		},
		Action: tradeRunner.Run,
	}

	return cmd
}
