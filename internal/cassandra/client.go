package cassandra

import (
	"time"

	"github.com/cpurta/tatanka/internal/model"
)

// Client is the interface that provides operations on a
// Cassandra cluster to read and write
// Trades, Periods, and Simulation results.
type Client interface {
	GetTradesBetween(start, end time.Time) ([]*model.Trade, error)
	InsertTrade(selector string, trade *model.Trade) error
	// GetResumeMarkers()
	// GetBalances()
	// GetSessions()
	// GetPeriods()
	// GetMyTrades()
	// GetSimResults()
}
