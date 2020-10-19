package exchanges

import (
	"net/http"

	coinbasepro "github.com/cpurta/tatanka/extensions/exchanges/coinbase-pro"
	"github.com/cpurta/tatanka/internal/config"
	"github.com/cpurta/tatanka/internal/model"
)

func GetExchange(exchangeID string, config *config.Config) (model.Exchange, error) {
	var (
		exchange model.Exchange
		err      error
	)

	switch exchangeID {
	case "coinbasepro", "gdax":
		if exchange, err = coinbasepro.NewCoinbaseProExchange(
			config.GDAXConfig.APIKey,
			config.GDAXConfig.APIPassphrase,
			config.GDAXConfig.APISecret,
			http.DefaultClient,
		); err != nil {
			return nil, err
		}

		return exchange, nil
	}

	return nil, nil
}
