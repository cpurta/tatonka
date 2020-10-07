package factories

import (
	"reflect"
	"strings"

	"github.com/cpurta/tatanka/extensions/strategies/macd"
	"github.com/cpurta/tatanka/extensions/strategies/rsi"
	"github.com/cpurta/tatanka/internal/config"
	"github.com/cpurta/tatanka/internal/model"
)

type strategyFactory struct{}

func NewStrategyFactory() *strategyFactory {
	return &strategyFactory{}
}

func (factory *strategyFactory) GetStrategies(config *config.Config) []model.Strategy {
	var (
		strategies = make([]model.Strategy, 0)
	)

	for _, strategy := range config.Strategies {
		switch strings.ToLower(strategy.Name) {
		case "macd":
			var (
				longValue  int64
				shortValue int64
			)

			if long, ok := strategy.Options["long"]; ok {
				longValue = reflect.ValueOf(long).Int()
			}

			if short, ok := strategy.Options["short"]; ok {
				shortValue = reflect.ValueOf(short).Int()
			}

			strategies = append(strategies, macd.NewMACD(int(shortValue), int(longValue)))
		case "rsi":
			strategies = append(strategies, rsi.NewRSI())
		}
	}

	return strategies
}
