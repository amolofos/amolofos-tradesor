package exporters

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/amolofos/tradesor/pkg/interfaces/canonical_models"
	"github.com/amolofos/tradesor/pkg/models/models_outputFormat"
)

type FileExporter struct {
	outputTo string
}

func NewFileExporter(outputTo string) (e *FileExporter, err error) {
	e = &FileExporter{}
	err = e.Init(outputTo)
	return
}

func (e *FileExporter) Init(outputTo string) (err error) {
	e.outputTo = outputTo
	return
}

func (e *FileExporter) Write(doc canonical_models.CanonicalModel, outputFormat models_outputFormat.OutputFormat) (nProducts int, err error) {
	nCategories := 0

	categories := doc.Categories()
	slog.Debug(fmt.Sprintf("FileExporter: categories %s", categories))

	for _, category := range categories {
		categoryFile := filepath.Join(e.outputTo, category+"."+outputFormat.String())

		nProductsCategory, productsExported, errProducts := doc.Export(category, outputFormat)
		if errProducts != nil {
			slog.Error(fmt.Sprintf("FileExporter: Error getting products for category: %s with error: %s", category, errProducts))
			err = errProducts

			// Let's create the other files before failing.
			continue
		}

		fo, errFileOut := os.Create(categoryFile)
		if errFileOut != nil {
			slog.Error(fmt.Sprintf("FileExporter: Error creating category file %s with error: %s", categoryFile, errFileOut))
			err = errFileOut

			// Let's create the other files before failing.
			continue
		}
		defer fo.Close()

		_, errWriteProducts := fo.WriteString(productsExported)
		if errWriteProducts != nil {
			slog.Error(fmt.Sprintf("FileExporter: Error writing products into file %s with error: %s", categoryFile, errWriteProducts))
			err = errWriteProducts
		}

		nProducts += nProductsCategory
		nCategories += 1
	}

	slog.Info(fmt.Sprintf("FileExporter: Created %d csv files with %d products in total.", nCategories, nProducts))
	return
}
