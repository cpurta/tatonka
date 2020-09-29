package exchanges

import (
	"net/http"

	"github.com/cpurta/tatanka/extensions/exchanges/gdax"
	"github.com/cpurta/tatanka/internal/config"
	"github.com/cpurta/tatanka/internal/model"
)

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
