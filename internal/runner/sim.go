package runner

import (
	"errors"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/cpurta/tatanka/internal/cassandra"
	"github.com/cpurta/tatanka/internal/config"
	"github.com/cpurta/tatanka/internal/factories"
	"github.com/cpurta/tatanka/internal/model"
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
	RSIPeriods           int64
	ExactBuyOrders       bool
	ExactSellOrders      bool
	DisableOptions       bool
	QuarentineTime       int64
	EnableStats          bool
	BacktesterGeneration int64
	Period               time.Duration
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
		endTime          = time.Now()
		trades           []*model.Trade
		periods          []*model.Period
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

	startTime := endTime.Add(time.Hour * -24 * time.Duration(runner.Days))

	if trades, err = runner.cassandraClient.GetTradesBetween(selector.String(), startTime, endTime); err != nil {
		return fmt.Errorf("unable to get trades from cassandra: %s", err.Error())
	}

	periods = runner.getTradePeriods(trades)

	for _, period := range periods {
		fmt.Println(period)
	}

	return nil
}

func (runner *SimRunner) getTradePeriods(trades []*model.Trade) []*model.Period {
	var (
		periodFactory     = factories.NewPeriodFactory()
		beginTradeTime    = trades[0].Time.Truncate(time.Minute)
		endTradeTime      = beginTradeTime.Add(runner.Period)
		tradePeriodBucket = make([]*model.Trade, 0)
		periods           = make([]*model.Period, 0)
	)

	for _, trade := range trades {
		if trade.Time.Before(endTradeTime) {
			tradePeriodBucket = append(tradePeriodBucket, trade)
			continue
		}

		period := periodFactory.GetPeriod(tradePeriodBucket)
		periods = append(periods, period)

		beginTradeTime = trade.Time.Truncate(time.Minute)
		endTradeTime = beginTradeTime.Add(runner.Period)
		tradePeriodBucket = tradePeriodBucket[:0]
		tradePeriodBucket = append(tradePeriodBucket, trade)
	}

	return periods
}
