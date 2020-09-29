package cassandra

import (
	"time"

	"github.com/cpurta/tatanka/internal/model"
)

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
