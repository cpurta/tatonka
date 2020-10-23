package exchanges

import (
	"net/http"

	coinbasepro "github.com/cpurta/tatanka/extensions/exchanges/coinbase-pro"
	"github.com/cpurta/tatanka/internal/config"
	"github.com/cpurta/tatanka/internal/model"
)

// GetExchange returns a Exchange model using the provided configuration.
// The concrete Exchange implementation is determined by the exchangeID, and
// the caller must be wary that this function will return nil if the exchangeID
// is unknown.
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
