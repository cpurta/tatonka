package rsi

import (
	"math"

	"github.com/cpurta/tatonka/internal/model"
)

var _ model.Strategy = &rsi{}

type rsi struct{}

func RSI() *rsi {
	return &rsi{}
}

func (r *rsi) Name() string {
	return "rsi"
}

func (r *rsi) Description() string {
	return "Attempts to buy low and sell high by tracking RSI high-water readings."
}

func (r *rsi) Options() []model.Option {
	return []model.Option{
		&model.IntOption{
			Name:         "rsi_periods",
			Description:  "number of RSI periods",
			DefaultValue: 14,
		},
		&model.IntOption{
			Name:         "oversold_rsi",
			Description:  "buy when RSI reaches or drops below this value",
			DefaultValue: 30,
		},
		&model.IntOption{
			Name:         "overbought_rsi",
			Description:  "sell when RSI reaches or goes above this value",
			DefaultValue: 82,
		},
		&model.IntOption{
			Name:         "rsi_recover",
			Description:  "allow RSI to recover this many points before buying",
			DefaultValue: 3,
		},
		&model.IntOption{
			Name:         "rsi_drop",
			Description:  "allow RSI to fall this many points before selling",
			DefaultValue: 0,
		},
		&model.IntOption{
			Name:         "rsi_divisor",
			Description:  "sell when RSI reaches high-water reading divided by this value",
			DefaultValue: 2,
		},
	}
}

func (r *rsi) Calculate(periods []*model.Period) (float64, error) {
	var (
		totalUp   = 0.0
		totalDown = 0.0
	)

	for i := 1; i < len(periods); i++ {
		currentClose := periods[i].Close
		previousClose := periods[i-1].Close

		closeDifference := currentClose - previousClose

		if closeDifference >= 0 {
			totalUp += closeDifference
		} else {
			totalDown -= closeDifference
		}
	}

	averageUp := totalUp / float64(len(periods))
	averageDown := totalDown / float64(len(periods))

	relativeStrength := averageUp / math.Abs(averageDown)

	return 100 - (100 / (1 + relativeStrength)), nil
}

func (r *rsi) Signal(rsi float64) model.Signal {
	if rsi < 30.0 {
		return model.BuySignal
	}

	if rsi > 82.0 {
		return model.SellSignal
	}

	return model.NeutralSignal
}
