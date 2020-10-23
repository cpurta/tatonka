package runner

import (
	"fmt"

	"github.com/cpurta/tatanka/extensions/strategies/macd"
	"github.com/cpurta/tatanka/extensions/strategies/rsi"
	"github.com/cpurta/tatanka/internal/model"
	. "github.com/logrusorgru/aurora"
	"github.com/urfave/cli/v2"
)

type ListStrategiesRunner struct{}

func (runner *ListStrategiesRunner) Run(cli *cli.Context) error {
	strategies := runner.getStrategies()

	for _, strategy := range strategies {
		fmt.Printf("%s\n", Cyan(strategy.Name()))
		fmt.Println(Sprintf("    %s", BrightBlack("description:")))
		fmt.Println(Sprintf("        %s", BrightBlack(strategy.Description())))
		fmt.Println(Sprintf("    %s", BrightBlack("options:")))
		for _, option := range strategy.Options() {
			fmt.Printf("        %s\n", option)
		}
	}

	return nil
}

func (runner *ListStrategiesRunner) getStrategies() []model.Strategy {
	return []model.Strategy{
		rsi.NewRSI(),
		macd.NewMACD(26, 12),
	}
}
