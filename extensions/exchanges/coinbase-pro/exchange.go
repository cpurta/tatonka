package coinbasepro

import (
	"net/http"
	"strconv"
	"sync"

	"github.com/cpurta/tatanka/internal/model"
	"github.com/gorilla/websocket"
	coinbasepro "github.com/preichenberger/go-coinbasepro/v2"
)

var (
	_ model.Exchange = &coinbaseProExchange{}
)

type coinbaseProExchange struct {
	client            *coinbasepro.Client
	wsConn            *websocket.Conn
	tradeChan         chan *model.Trade
	tradeCursor       *coinbasepro.Cursor
	historyScan       string
	makerFee          float64
	takerFee          float64
	BackfillRateLimit int64
}

func NewCoinbaseProExchange(key, passphrase, secret string, httpClient *http.Client) (*coinbaseProExchange, error) {
	var (
		client = coinbasepro.NewClient()
		dialer websocket.Dialer
		conn   *websocket.Conn
		err    error
	)

	client.UpdateConfig(&coinbasepro.ClientConfig{
		BaseURL:    "https://api.pro.coinbase.com",
		Key:        key,
		Passphrase: passphrase,
		Secret:     secret,
	})

	client.HTTPClient = httpClient

	if conn, _, err = dialer.Dial("wss://ws-feed.pro.coinbase.com", nil); err != nil {
		return nil, err
	}

	return &coinbaseProExchange{
		client:            client,
		wsConn:            conn,
		tradeChan:         make(chan *model.Trade, 1000),
		historyScan:       "backward",
		makerFee:          0.0,
		takerFee:          0.3,
		BackfillRateLimit: int64(335),
	}, nil
}

func (exchange *coinbaseProExchange) Start(productID string) error {
	var once sync.Once

	start := func() {
		var (
			subscribe = coinbasepro.Message{
				Type: "subscribe",
				Channels: []coinbasepro.MessageChannel{
					coinbasepro.MessageChannel{
						Name: "heartbeat",
						ProductIds: []string{
							productID,
						},
					},
					coinbasepro.MessageChannel{
						Name: "level2",
						ProductIds: []string{
							productID,
						},
					},
				},
			}
			err error
		)

		if err = exchange.wsConn.WriteJSON(subscribe); err != nil {
			println("unable to write json message to coinbase pro websocket", err.Error())
		}
	}

	once.Do(start)

	for {
		message := coinbasepro.Message{}

		if err := exchange.wsConn.ReadJSON(&message); err != nil {
			println("unable to read json message from coinbase pro websocket feed", err.Error())
			continue
		}

		println("recieved coinbase pro message:", message.Type)
	}

	return nil
}

func (exchange *coinbaseProExchange) Trades(productID string) <-chan *model.Trade {
	return exchange.tradeChan
}

func (exchange *coinbaseProExchange) HistoricalTrades(productID string) ([]*model.Trade, error) {
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
		size, _ := strconv.ParseFloat(trade.Size, 64)

		price, _ := strconv.ParseFloat(trade.Price, 64)

		trades = append(trades, &model.Trade{
			TradeID: strconv.Itoa(trade.TradeID),
			Size:    size,
			Price:   price,
			Time:    trade.Time.Time(),
			Side:    trade.Side,
		})
	}

	return trades, nil
}

func (exchange *coinbaseProExchange) Balance(currency string, asset string) (*model.Balance, error) {
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

func (exchange *coinbaseProExchange) Quote(productID string) (*model.Quote, error) {
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
