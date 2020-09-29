package factories

import (
	"strconv"

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
		tradePrice, _ := strconv.ParseFloat(trade.Price, 64)

		period.High = max(period.High, tradePrice)
		period.Low = min(period.Low, tradePrice)

		if i == len(trades)-1 {
			period.Close = tradePrice
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
