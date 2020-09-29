package model

type Exchange interface {
	GetTrades(productID string) ([]*Trade, error)
	GetBalance(currency string, asset string) (*Balance, error)
	GetQuote(productID string) (*Quote, error)
}
