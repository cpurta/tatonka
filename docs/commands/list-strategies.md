```
rsi
    description:
        Attempts to buy low and sell high by tracking RSI high-water readings.
    options:
        --period=<value>   period length, same as --period_length (default: 2m0s)
        --period_length=<value>   period length, same as --period (default: 2m0s)
        --min_periods=<value>   min. number of history periods (default: 52)
        --rsi_periods=<value>   number of RSI periods (default: 14)
        --oversold_rsi=<value>   buy when RSI reaches or drops below this value (default: 30)
        --overbought_rsi=<value>   sell when RSI reaches or goes above this value (default: 82)
        --rsi_recover=<value>   allow RSI to recover this many points before buying (default: 3)
        --rsi_drop=<value>   allow RSI to fall this many points before selling (default: 0)
        --rsi_divisor=<value>   sell when RSI reaches high-water reading divided by this value (default: 2)
```
