package runner

import (
	"errors"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/cpurta/tatanka/internal/cassandra"
	"github.com/cpurta/tatanka/internal/config"
	"github.com/cpurta/tatanka/internal/exchanges"
	"github.com/cpurta/tatanka/internal/model"
	"github.com/gocql/gocql"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v2"
)

// TradeRunner holds the necessary data an implementation for the
// 'trade' CLI command. This runner is able to connect to a Cassandra
// cluster (with the user-provided configuration) and fetch/insert
// trading data.
type TradeRunner struct {
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
	PaperTrade           bool
	cassandraClient      cassandra.Client
}

func (runner *TradeRunner) Run(cli *cli.Context) error {
	var (
		configFile       []byte
		config           *config.Config
		selector         *model.Selector
		exchange         model.Exchange
		cassandraCluster *gocql.ClusterConfig
		cassandraSession *gocql.Session
		startBackfill    = time.Now().Add(time.Hour * -24).Truncate(time.Hour)
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

	exchange = exchanges.GetExchange(selector.ExchangeID, config)

	if exchange == nil {
		return fmt.Errorf("%s exchange is not supported", selector.ExchangeID)
	}

	println("fetching pre-roll data:")

	for {
		trades, err := exchange.GetTrades(selector.ProductID)
		if err != nil {
			return fmt.Errorf("unable to get historical trades for %s: %s", selector.ProductID, err.Error())
		}

		if len(trades) == 0 {
			println("recieved 0 historical trades from exchange, moving on to watching market...")
			break
		}

		if trades[0].Time.Before(startBackfill) {
			println("finished backfilling pre-roll data")
			break
		}

		runner.insertTrades(selector.String(), trades)
	}

	if runner.PaperTrade {
		println("----- STARTING PAPER TRADING -----")
	} else {
		println("----- STARTING LIVE TRADING ------")
	}

	for {
		println("mock watching live market data")

		// TODO: watch live market data and feed to strategies to determine buy/sell signals

		time.Sleep(time.Minute * 1)
	}

	return nil
}

func (runner *TradeRunner) insertTrades(selector string, trades []*model.Trade) {
	for _, trade := range trades {
		if err := runner.cassandraClient.InsertTrade(selector, trade); err != nil {
			println("error inserting trades into cassandra:", err.Error())
		}
	}
}
