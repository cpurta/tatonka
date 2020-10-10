package model

// Exchange provides methods to get trades, balances, and product id quotes from
// a supported crypto-currency exchange
type Exchange interface {
	GetTrades(productID string) ([]*Trade, error)
	GetBalance(currency string, asset string) (*Balance, error)
	GetQuote(productID string) (*Quote, error)
}
