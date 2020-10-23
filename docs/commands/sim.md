```
NAME:
   tatanka sim - run a simulation on backfilled data

USAGE:
   tatanka sim [command options] [arguments...]

OPTIONS:
   --config value                  path to optional config overrides file (default: "/etc/tatanka/config.yaml")
   --strategy value                strategy to use (default: "trend_ema")
   --order_type value              order type to use (maker/taker) (default: "maker")
   --filename value                filename for the result output (ex: result.html). "none" to disable
   --days value                    set duration by day count (default: 14)
   --currency_capital value        amount of start capital in currency (default: 1000)
   --asset_capital value           amount of start capital in asset (default: 0)
   --avg_slippage_pct value        avg. amount of slippage to apply to trades (default: 0.045)
   --buy_pct value                 buy with this % of currency balance (default: 99)
   --sell_pct value                sell with this % of asset balance (default: 99)
   --markdown_buy_pct value        % to mark down buy price (default: 0)
   --markdown_sell_pct value       % to mark up sell price (default: 0)
   --order_adjust_time value       adjust bid/ask on this interval to keep orders competitive in ms (default: 5000)
   --order_poll_time value         poll order status on this interval in ms (default: 5000)
   --sell_cancel_pct value         cancels the sale if the price is between this percentage (for more or less) (default: 0)
   --sell_stop_pct value           sell if price drops below this % of bought price (default: 0)
   --buy_stop_pct value            buy if price surges above this % of sold price (default: 0)
   --profit_stop_enable_pct value  enable trailing sell stop when reaching this % profit (default: 0)
   --profit_stop_pct value         maintain a trailing stop this % below the high-water mark of profit (default: 1)
   --max_sell_loss_pct value       avoid selling at a loss pct under this float (default: 99)
   --max_buy_loss_pct value        avoid buying at a loss pct over this float (default: 99)
   --max_slippage_pct value        avoid selling at a slippage pct above this float (default: 5)
   --rsi_periods value             number of periods to calculate RSI at (default: 14)
   --exact_buy_orders              instead of only adjusting maker buy when the price goes up, adjust it if price has changed at all (default: false)
   --exact_sell_orders             instead of only adjusting maker sell when the price goes down, adjust it if price has changed at all (default: false)
   --disable_options               disable printing of options (default: false)
   --quarentine_time value         For loss trade, set quarentine time for cancel buys in minutes (default: 10)
   --enable_stats                  enable printing order stats (default: false)
   --backtester_generation value   creates a json file in simulations with the generation number (default: -1)
   --verbose                       print status lines on every period (default: false)
   --silent                        only output on completion (can speed up sim) (default: false)
   --help, -h                      show help (default: false)
```
