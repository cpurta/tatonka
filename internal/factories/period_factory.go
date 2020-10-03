package factories

import (
	"strconv"

	"github.com/cpurta/tatanka/internal/model"
)

// PeriodFactory is a factory of Periods.
// Each Period consolidates data about a range of changes, and can be
// obtained through the  GetPeriod method.
type PeriodFactory struct{}

// NewPeriodFactory returns a new instance of PeriodFactory.
func NewPeriodFactory() *PeriodFactory {
	return &PeriodFactory{}
}

// GetPeriod consolidates information about a certain amount of trades, returning
// the high, low and closing value of the period.
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
