package runner

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/cpurta/tatanka/internal/model"
	. "github.com/logrusorgru/aurora"
	"github.com/urfave/cli/v2"
)

type ListSelectorRunner struct{}

func (runner *ListSelectorRunner) Run(cli *cli.Context) error {
	var (
		files []os.FileInfo
		err   error
	)

	if files, err = ioutil.ReadDir("./extensions/exchanges"); err != nil {
		return nil
	}

	for _, exchangeName := range files {
		var (
			productFile     []byte
			products        []*model.Product
			productFileName = fmt.Sprintf("./extensions/exchanges/%s/products.json", exchangeName.Name())
		)
		fmt.Printf("%s:\n", exchangeName.Name())

		if productFile, err = ioutil.ReadFile(productFileName); err != nil {
			fmt.Printf("unable to read %s: %s\n", productFileName, err.Error())
			continue
		}

		if err = json.Unmarshal(productFile, &products); err != nil {
			fmt.Println("unable to unmarshal products file:", err.Error())
			continue
		}

		for _, product := range products {
			fmt.Println(Sprintf("  %s.%s-%s   %s%s%s", Cyan(exchangeName.Name()), Green(product.Asset), Cyan(product.Currency), BrightBlack("("), BrightBlack(product.Label), BrightBlack(")")))
		}
	}

	return nil
}
