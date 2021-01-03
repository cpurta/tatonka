package runner

import (
	"errors"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/cpurta/tatonka/internal/config"
	"github.com/cpurta/tatonka/internal/exchanges"
	"github.com/cpurta/tatonka/internal/model"
	. "github.com/logrusorgru/aurora"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v2"
)

const dateTimeFormat = "2006-01-02 15:04:05"

type BalanceRunner struct {
	ConfigFile        string
	CalculateCurrency bool
}

func (runner *BalanceRunner) Run(cli *cli.Context) error {
	var (
		configFile []byte
		config     *config.Config
		selector   *model.Selector
		exchange   model.Exchange
		balance    *model.Balance
		quote      *model.Quote
		now        = time.Now().UTC()
		err        error
	)

	if cli.NArg() == 0 {
		return errors.New("you must specify a selector (i.e. {exchange_slug}.{product_id})")
	}

	if selector, err = model.NewSelectorFromString(cli.Args().Get(0)); err != nil {
		return err
	}

	if configFile, err = ioutil.ReadFile(runner.ConfigFile); err != nil {
		return err
	}

	if err = yaml.Unmarshal(configFile, &config); err != nil {
		return err
	}

	exchange = exchanges.GetExchange(selector.ExchangeID, config)

	if exchange == nil {
		return fmt.Errorf("%s exchange is not supported", selector.ExchangeID)
	}

	if balance, err = exchange.GetBalance(selector.Currency, selector.Asset); err != nil {
		return fmt.Errorf("unable to get balance from exchange: %s", err.Error())
	}

	if quote, err = exchange.GetQuote(selector.ProductID); err != nil {
		return fmt.Errorf("unable to get quote from exchange: %s", err.Error())
	}

	println(Sprintf("%s %.2f %s", BrightBlack(now.Format(dateTimeFormat)), Cyan(quote.Ask), BrightBlack(selector.ProductID)))
	println(Sprintf("%s %s %.8f %s %.8f", BrightBlack(now.Format(dateTimeFormat)), BrightBlack("Asset:"), balance.Asset, BrightBlack("Available:"), Yellow(balance.Asset-balance.AssetHold)))
	println(Sprintf("%s %s %.8f", BrightBlack(now.Format(dateTimeFormat)), BrightBlack("Asset Value:"), balance.Asset*quote.Ask))
	println(Sprintf("%s %s %.8f %s %.8f", BrightBlack(now.Format(dateTimeFormat)), BrightBlack("Currency:"), balance.Currency, BrightBlack("Available:"), Yellow(balance.Currency-balance.CurrencyHold)))
	println(Sprintf("%s %s %.8f", BrightBlack(now.Format(dateTimeFormat)), BrightBlack("Total:"), balance.Asset*quote.Ask+balance.Currency))

	return nil
}
