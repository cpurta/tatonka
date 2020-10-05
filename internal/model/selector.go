package model

import (
	"errors"
	"fmt"
	"strings"
)

// Selector holds info related to the exchange, and asset/currency pair tatanka
// should pull trade info, balance info for
type Selector struct {
	ExchangeID string
	ProductID  string
	Asset      string
	Currency   string
}

// String provides a string representaion of a Selector and will return in the
// format "{exchange_slug}.{product_id}" (e.g. gdax.BTC-USD)
func (selector *Selector) String() string {
	return fmt.Sprintf("%s.%s", selector.ExchangeID, selector.ProductID)
}

func NewSelectorFromString(selectorStr string) (*Selector, error) {
	var (
		parts      []string
		pairParts  []string
		exchangeID string
		productID  string
		asset      string
		currency   string
	)

	parts = strings.Split(selectorStr, ".")

	if len(parts) == 0 {
		return nil, errors.New("selector argument should be in format: {exchange_slug}.{product_id} (e.g. gdax.BTC-USD)")
	}

	exchangeID = parts[0]
	productID = parts[1]

	if !strings.Contains(productID, "-") {
		return nil, errors.New("product_id specified must contain dash between currency pairs (e.g. gdax.BTC-USD)")
	}

	pairParts = strings.Split(productID, "-")

	asset = pairParts[0]
	currency = pairParts[1]

	return &Selector{
		ExchangeID: exchangeID,
		ProductID:  productID,
		Asset:      asset,
		Currency:   currency,
	}, nil
}
