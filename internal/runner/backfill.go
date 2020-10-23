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

type BackfillRunner struct {
	ConfigFile      string
	Days            int64
	Debug           bool
	cassandraClient cassandra.Client
}

func (runner *BackfillRunner) Run(cli *cli.Context) error {
	var (
		configFile       []byte
		config           *config.Config
		selector         *model.Selector
		exchange         model.Exchange
		cassandraCluster *gocql.ClusterConfig
		cassandraSession *gocql.Session
		err              error
	)

	startBackfill := time.Now().Add(time.Hour * -24 * time.Duration(runner.Days)).Truncate(time.Hour)

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

	if exchange, err = exchanges.GetExchange(selector.ExchangeID, config); err != nil {
		return fmt.Errorf("%s exchange is not supported", selector.ExchangeID)
	}

	fmt.Printf("backfilling %d of days historical data for %s\n", runner.Days, selector.String())

	for {
		trades, err := exchange.HistoricalTrades(selector.ProductID)
		if err != nil {
			return fmt.Errorf("unable to get historical trades for %s: %s", selector.ProductID, err.Error())
		}

		if len(trades) == 0 {
			break
		}

		if trades[0].Time.Before(startBackfill) {
			break
		}

		runner.insertTrades(selector.String(), trades)

		print(".")
	}

	return nil
}

func (runner *BackfillRunner) insertTrades(selector string, trades []*model.Trade) {
	for _, trade := range trades {
		if err := runner.cassandraClient.InsertTrade(selector, trade); err != nil {
			println("error inserting trades into cassandra:", err.Error())
		}
	}
}
