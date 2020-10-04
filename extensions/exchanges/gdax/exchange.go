package gdax

import (
	"net/http"
	"strconv"

	"github.com/cpurta/tatanka/internal/model"
	coinbasepro "github.com/preichenberger/go-coinbasepro/v2"
)

var (
	_ model.Exchange = &gdaxExchange{}
)

// gdaxExchange implements the representation of a coinbase(gdax) exchange
type gdaxExchange struct {
	client            *coinbasepro.Client
	tradeCursor       *coinbasepro.Cursor
	historyScan       string
	makerFee          float64
	takerFee          float64
	BackfillRateLimit int64
}

// NewGDAXExchange add new exchange to coinbase(gdax)
func NewGDAXExchange(key, passphrase, secret string, httpClient *http.Client) *gdaxExchange {
	client := coinbasepro.NewClient()

	client.UpdateConfig(&coinbasepro.ClientConfig{
		BaseURL:    "https://api.pro.coinbase.com",
		Key:        key,
		Passphrase: passphrase,
		Secret:     secret,
	})

	client.HTTPClient = httpClient

	return &gdaxExchange{
		client:            client,
		historyScan:       "backward",
		makerFee:          0.0,
		takerFee:          0.3,
		BackfillRateLimit: int64(335),
	}
}

// GetTrades  returns coinbase(gdax) transaction history
func (exchange *gdaxExchange) GetTrades(productID string) ([]*model.Trade, error) {
	var (
		gdaxTrades []coinbasepro.Trade
		trades     = make([]*model.Trade, 0)
	)

	if exchange.tradeCursor == nil {
		exchange.tradeCursor = exchange.client.ListTrades(productID)
	}

	if !exchange.tradeCursor.HasMore {
		return trades, nil
	}

	if err := exchange.tradeCursor.NextPage(&gdaxTrades); err != nil {
		return nil, err
	}

	for _, trade := range gdaxTrades {
		trades = append(trades, &model.Trade{
			TradeID: strconv.Itoa(trade.TradeID),
			Size:    trade.Size,
			Price:   trade.Price,
			Time:    trade.Time.Time(),
			Side:    trade.Side,
		})
	}

	return trades, nil
}

// GetBalance return account balance on coinbase(gdax)
func (exchange *gdaxExchange) GetBalance(currency string, asset string) (*model.Balance, error) {
	var (
		accounts []coinbasepro.Account
		balance  = &model.Balance{}
		err      error
	)

	if accounts, err = exchange.client.GetAccounts(); err != nil {
		return nil, err
	}

	for _, account := range accounts {
		if account.Currency == currency {
			balance.Currency, _ = strconv.ParseFloat(account.Balance, 64)
			balance.CurrencyHold, _ = strconv.ParseFloat(account.Hold, 64)
		}

		if account.Currency == asset {
			balance.Asset, _ = strconv.ParseFloat(account.Balance, 64)
			balance.AssetHold, _ = strconv.ParseFloat(account.Hold, 64)
		}
	}

	return balance, nil
}

// GetQuote returns the price for the exchange made
func (exchange *gdaxExchange) GetQuote(productID string) (*model.Quote, error) {
	var (
		ticker coinbasepro.Ticker
		quote  = &model.Quote{}
		err    error
	)

	if ticker, err = exchange.client.GetTicker(productID); err != nil {
		return nil, err
	}

	quote.Bid, _ = strconv.ParseFloat(ticker.Bid, 64)
	quote.Ask, _ = strconv.ParseFloat(ticker.Ask, 64)

	return quote, nil
}
