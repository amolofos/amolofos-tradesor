package exporters

import (
	"github.com/amolofos/tradesor/pkg/interfaces/canonical_models"
	"github.com/amolofos/tradesor/pkg/models/models_outputFormat"
)

type HttpExporter struct{}

func NewHttpExporter(outputTo string) (e *HttpExporter, err error) {
	e = &HttpExporter{}
	err = e.Init(outputTo)
	return
}

func (e *HttpExporter) Init(outputTo string) (err error) {
	return
}

func (e *HttpExporter) Write(doc canonical_models.CanonicalModel, outputFormat models_outputFormat.OutputFormat) (nProducts int, err error) {
	return
}
