package exporters

import (
	"github.com/amolofos/tradesor/pkg/interfaces/canonical_models"
	"github.com/amolofos/tradesor/pkg/models/models_outputFormat"
)

type FileExporter struct{}

func NewFileExporter(outputTo string) (e *FileExporter, err error) {
	e = &FileExporter{}
	err = e.Init(outputTo)
	return
}

func (e *FileExporter) Init(outputTo string) (err error) {
	return
}

func (e *FileExporter) Write(doc canonical_models.CanonicalModel, outputFormat models_outputFormat.OutputFormat) (nProducts int, err error) {
	return
}
