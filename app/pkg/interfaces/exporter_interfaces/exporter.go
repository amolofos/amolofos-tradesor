package exporter_interfaces

import (
	"github.com/amolofos/tradesor/pkg/interfaces/canonical_models"
	"github.com/amolofos/tradesor/pkg/models/models_outputFormat"
)

type Exporter interface {
	Init(outputTo string) (err error)
	Write(doc canonical_models.CanonicalModel, outputFormat models_outputFormat.OutputFormat) (nProducts int, err error)
}
