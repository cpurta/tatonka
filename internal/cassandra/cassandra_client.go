package cassandra

import (
	"time"

	"github.com/cpurta/tatanka/internal/model"
	"github.com/gocql/gocql"
)

var _ Client = &cassandraClient{}

type cassandraClient struct {
	session *gocql.Session
}

// NewCassandraClient returns a new Cassandra Client interface for interacting with Cassandra
func NewCassandraClient(session *gocql.Session) *cassandraClient {
	return &cassandraClient{
		session: session,
	}
}

// GetTradesBetween returns the transactions that occurred in a date range
func (client *cassandraClient) GetTradesBetween(start, end time.Time) ([]*model.Trade, error) {
	var (
		query  = `SELECT trade_id, price, size, time, side FROM trades WHERE selector = ? AND time BETWEEN ? AND ?`
		iter   *gocql.Iter
		trade  = &model.Trade{}
		trades = make([]*model.Trade, 0)
		err    error
	)

	iter = client.session.Query(query, start.Format(time.RFC3339), end.Format(time.RFC3339)).Iter()

	for iter.Scan(&trade.TradeID, &trade.Price, &trade.Size, &trade.Time, &trade.Side) {
		trades = append(trades, trade)
	}

	if err = iter.Close(); err != nil {
		return trades, err
	}

	return trades, nil
}

// InsertTrade add new trade
func (client *cassandraClient) InsertTrade(selector string, trade *model.Trade) error {
	var (
		query = `INSERT INTO trades (id,selector,trade_id,price,size,time,side) VALUES (?,?,?,?,?,?,?)`
	)

	return client.session.Query(query,
		gocql.TimeUUID(),
		selector,
		trade.TradeID,
		trade.Price,
		trade.Size,
		trade.Time.Format(time.RFC3339),
		trade.Side,
	).Exec()
}
