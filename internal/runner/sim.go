package runner

import (
	"github.com/cpurta/tatanka/internal/cassandra"
	"github.com/urfave/cli/v2"
)

type SimRunner struct {
	ConfigFile           string
	Strategy             string
	OrderType            string
	Filename             string
	Days                 int64
	CurrencyCapital      int64
	AssetCapital         int64
	AvgSlippagePct       float64
	BuyPct               int64
	SellPct              int64
	MarkdownBuyPct       int64
	MarkupSellPct        int64
	OrderAdjustTime      int64
	OrderPollTime        int64
	SellCancelPct        float64
	SellStopPct          float64
	BuyStopPct           float64
	ProfitStopEnablePct  int64
	ProfitStopPct        int64
	MaxSellLossPct       int64
	MaxBuyLossPct        int64
	MaxSlippagePct       int64
	symmetrical          bool
	RSIPeriods           int64
	ExactBuyOrders       bool
	ExactSellOrders      bool
	DisableOptions       bool
	QuarentineTime       int64
	EnableStats          bool
	BacktesterGeneration int64
	Verbose              bool
	Silent               bool
	cassandraClient      cassandra.Client
}

func (runner *SimRunner) Run(cli *cli.Context) error {
	return nil
}
