package canonical_models

import (
	"github.com/amolofos/tradesor/pkg/models/models_outputFormat"
)

type CanonicalModel interface {
	Init()
	GetHeader() []string
	GetCategories() []string
	GetProductsFormatted(format models_outputFormat.OutputFormat) (output string, err error)
	GetProductsForCategory(category string) (products [][]string)
	GetProductsForCategoryFormatted(format models_outputFormat.OutputFormat, category string) (output string, err error)
}
