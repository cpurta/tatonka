package factories

import (
	"github.com/cpurta/tatanka/internal/model"
)

type PeriodFactory struct{}

func NewPeriodFactory() *PeriodFactory {
	return &PeriodFactory{}
}

func (factory *PeriodFactory) GetPeriod(trades model.Trades) *model.Period {
	var (
		period = &model.Period{}
	)

	for i, trade := range trades {
		period.High = max(period.High, trade.Price)
		period.Low = min(period.Low, trade.Price)

		if i == len(trades)-1 {
			period.Close = trade.Price
		}
	}

	return period
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}

	return b
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}

	return b
}
