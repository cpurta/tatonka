package model

// Exchange provides methods to get trades, balances, and product id quotes from
// a supported crypto-currency exchange
type Exchange interface {
	Trades(productID string) <-chan *Trade
	HistoricalTrades(productID string) ([]*Trade, error)
	Balance(currency string, asset string) (*Balance, error)
	Quote(productID string) (*Quote, error)
	Start(productID string) error
}
