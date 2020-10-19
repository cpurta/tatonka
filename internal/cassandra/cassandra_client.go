package cassandra

import (
	"fmt"
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
func (client *cassandraClient) GetTradesBetween(productID string, start, end time.Time) ([]*model.Trade, error) {
	var (
		query  = `SELECT trade_id, price, size, time, side FROM trades WHERE selector = ? AND (time >= ? AND time <= ?)`
		iter   *gocql.Iter
		trade  = &model.Trade{}
		trades = make([]*model.Trade, 0)
		err    error
	)

	iter = client.session.Query(query, productID, start.Format(time.RFC3339), end.Format(time.RFC3339)).Iter()

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
		query   = `INSERT INTO trades (id,selector,trade_id,price,size,time,side) VALUES (?,?,?,?,?,?,?)`
		tradeID = trade.ID()
		uuid    gocql.UUID
		err     error
	)

	if len(tradeID) != 16 {
		return fmt.Errorf("trade id must be exaclty 16 bytes long: %s", tradeID)
	}

	if uuid, err = gocql.UUIDFromBytes([]byte(tradeID)); err != nil {
		return fmt.Errorf("unable to create gocql UUID: %s", err.Error())
	}

	return client.session.Query(query,
		uuid,
		selector,
		trade.TradeID,
		trade.Price,
		trade.Size,
		trade.Time,
		trade.Side,
	).Exec()
}
