package canonical_models

import (
	"github.com/amolofos/tradesor/pkg/models/models_outputFormat"
)

type CanonicalModel interface {
	Init()

	Header() []string
	Categories() []string
	ProductIds(category string) (productIds []string)
	Products(category string) (products [][]string)

	FormatProduct(product string, format models_outputFormat.OutputFormat) (output string, err error)
	FormatProducts(category string, format models_outputFormat.OutputFormat) (output string, err error)
}
