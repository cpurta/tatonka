package exchanges

import (
	"net/http"

	"github.com/cpurta/tatanka/extensions/exchanges/gdax"
	"github.com/cpurta/tatanka/internal/config"
	"github.com/cpurta/tatanka/internal/model"
)

// GetExchange returns a Exchange model using the provided configuration.
// The concrete Exchange implementation is determined by the exchangeID, and
// the caller must be wary that this function will return nil if the exchangeID
// is unknown.
func GetExchange(exchangeID string, config *config.Config) model.Exchange {
	switch exchangeID {
	case "gdax":
		return gdax.NewGDAXExchange(
			config.GDAXConfig.APIKey,
			config.GDAXConfig.APIPassphrase,
			config.GDAXConfig.APISecret,
			http.DefaultClient,
		)
	}

	return nil
}
