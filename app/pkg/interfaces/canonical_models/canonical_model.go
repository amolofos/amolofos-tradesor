package canonical_models

import (
	"github.com/amolofos/tradesor/pkg/models/models_outputFormat"
)

type CanonicalModel interface {
	Init()

	Header() []string
	Categories() []string
	ProductIds(category string) (productIds []string, err error)
	Products(category string) (products [][]string, err error)

	ExportHeader(outputFormat models_outputFormat.OutputFormat) (header string, err error)
	Export(category string, outputFormat models_outputFormat.OutputFormat) (nProducts int, products string, err error)
}
