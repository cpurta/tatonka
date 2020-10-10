package model

// Balance holds the account info related to an exchange account. This includes
// asset and base currency amounts and holds.
type Balance struct {
	Asset        float64
	AssetHold    float64
	Currency     float64
	CurrencyHold float64
}
