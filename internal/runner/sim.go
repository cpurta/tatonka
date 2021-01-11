package runner

import (
	"errors"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/cpurta/tatonka/extensions/strategies/rsi"
	"github.com/cpurta/tatonka/internal/cassandra"
	"github.com/cpurta/tatonka/internal/config"
	"github.com/cpurta/tatonka/internal/factories"
	"github.com/cpurta/tatonka/internal/model"
	"github.com/gocql/gocql"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v2"
)

// SimRunner implements the trading simulation logic for the
// 'sim' CLI command.
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
	PeriodDuration       time.Duration
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
	var (
		configFile       []byte
		config           *config.Config
		selector         *model.Selector
		cassandraCluster *gocql.ClusterConfig
		cassandraSession *gocql.Session
		now              = time.Now()
		since            = now.Add(time.Hour * -24 * time.Duration(runner.Days)).Truncate(time.Hour)
		trades           []*model.Trade
		periodTrades     = make([]*model.Trade, 0)
		periodFactory    = factories.NewPeriodFactory()
		periods          []*model.Period
		rsiStrategy      = rsi.RSI()
		signal           model.Signal
		rsiFloat float64
		err              error
	)

	if cli.NArg() == 0 {
		return errors.New("you must specify a selector (i.e. {exchange_slug}.{product_id})")
	}

	if selector, err = model.NewSelectorFromString(cli.Args().Get(0)); err != nil {
		return err
	}

	if configFile, err = ioutil.ReadFile(runner.ConfigFile); err != nil {
		return err
	}

	if err = yaml.Unmarshal(configFile, &config); err != nil {
		return err
	}

	cassandraCluster = gocql.NewCluster(config.CassandraConfig.Cluster...)
	cassandraCluster.Keyspace = config.CassandraConfig.Keyspace
	cassandraCluster.Consistency = gocql.Quorum

	if cassandraSession, err = cassandraCluster.CreateSession(); err != nil {
		return fmt.Errorf("unable to connect to cassandra cluster: %s", err.Error())
	}

	defer cassandraSession.Close()

	runner.cassandraClient = cassandra.NewCassandraClient(cassandraSession)

	if trades, err = runner.cassandraClient.GetTradesBetween(selector.String(), now, since); err != nil {
		return fmt.Errorf("unable to get historical trades: %s", err.Error())
	}

	var (
		periodStart = time.Now()
		periodEnd   = periodStart.Add(runner.PeriodDuration)
		period      *model.Period
	)

	for _, trade := range trades {
		if trade.Time.After(periodEnd) {
			period = periodFactory.GetPeriod(periodTrades)

			periods = append(periods, period)

			periodTrades = make([]*model.Trade, 0)

			periodStart.Add(runner.PeriodDuration)
			periodEnd.Add(runner.PeriodDuration)
		}

		if trade.Time.After(periodStart) && trade.Time.Before(periodEnd) {
			periodTrades = append(periodTrades, trade)
		}
	}

	if rsiFloat, err = rsiStrategy.Calculate(periods[:14]); err != nil {
		fmt.Println("unable to calculate rsi", err.Error())
	}

	signal = rsiStrategy.Signal()

	fmt.Println("RSI Signal:")

	for i := 1; i < len(periods)-14; i++ {
		if rsiFloat, err = rsiStrategy.Calculate(periods[i : i+14]); {
			fmt.Println("unable to calculate rsi", err.Error())
		}

		signal = rsiStrategy.Signal(rsiFloat)
	}

	return nil
}
