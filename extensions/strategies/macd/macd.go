package macd

import (
	"fmt"

	"github.com/cpurta/tatanka/internal/model"
)

var _ model.Strategy = &macd{}

type macd struct {
	short int
	long  int
}

func NewMACD(short, long int) *macd {
	return &macd{
		short: short,
		long:  long,
	}
}

func (m *macd) Name() string {
	return "macd"
}

func (m *macd) Description() string {
	return "Attempts to buy low and sell high by tracking MACD momentum readings."
}

func (m *macd) Options() []model.Option {
	return []model.Option{}
}

func (m *macd) Calculate(periods []*model.Period) (float64, error) {
	if len(periods) < m.long+1 {
		return 0.0, fmt.Errorf("must provided at least %d periods to calculate MACD indicator", m.long)
	}

	emaShort := m.ema(m.short, periods)
	emaLong := m.ema(m.long, periods)

	return emaShort - emaLong, nil
}

func (m *macd) ema(length int, periods []*model.Period) float64 {
	var (
		total           = 0.0
		smoothingFactor = 2.0 / (float64(length) + 1.0)
		emas            = make([]float64, len(periods)-length)
	)

	for i := 0; i < length; i++ {
		total += periods[i].Close
	}

	emas[0] = total / float64(length)

	j := 1
	for i := length; i < len(periods); i++ {
		emas[j] = periods[i].Close*smoothingFactor + (emas[j-1] * smoothingFactor)
		j++
	}

	return emas[len(emas)-1]
}

func (m *macd) Signal(macd float64) model.Signal {
	return model.SellSignal
}
