package service_exporter

import (
	"encoding/csv"
	"fmt"
	"log/slog"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/amolofos/tradesor/pkg/interfaces/canonical_models"
	"github.com/amolofos/tradesor/pkg/models/models_outputFormat"
)

type Exporter struct {
	replacer *strings.Replacer
}

func (u *Exporter) Init() {
	u.replacer = strings.NewReplacer(" ", "", "/", "", ">", "_")
}

func (u *Exporter) Export(doc canonical_models.CanonicalModel, outputFormat models_outputFormat.OutputFormat, outputTo string) (err error) {
	//remoteUrl, err := url.ParseRequestURI(outputTo)
	//if err != nil {
	//	u.exportToLocalDir(doc, outputTo)
	//	return
	//}
	//
	//u.exportToWordpress(doc, remoteUrl)
	u.createOutputDir(outputTo)
	u.exportToLocalDir(doc, outputFormat, outputTo)

	return
}

func (u *Exporter) exportToLocalDir(doc canonical_models.CanonicalModel, outputFormat models_outputFormat.OutputFormat, outputTo string) {
	categories := doc.GetCategories()
	for _, category := range categories {
		categoryFile := filepath.Join(outputTo, category+".csv")
		header := doc.GetHeader()

		data, errProducts := doc.GetProductsForCategoryFormatted(outputFormat, category)
		if errProducts != nil {
			slog.Error("Error getting products for category: ", category, " with error: ", errProducts)
		}

		fo, errFileOut := os.Create(categoryFile)
		if errFileOut != nil {
			slog.Error("Error creating category file: ", categoryFile, " with error: ", errFileOut)
		}
		defer fo.Close()

		csvWrite := csv.NewWriter(fo)
		csvWrite.Write(header)
		errCsvWrite := csvWrite.Error()
		if errCsvWrite != nil {
			slog.Error("Error writing category file: ", categoryFile, " with error: ", errCsvWrite)
		}

		fo.WriteString(data)
	}

	slog.Info(fmt.Sprintf("Created %d csv files.", len(categories)))
}

func (u *Exporter) exportToWordpress(doc canonical_models.CanonicalModel, remoteUrl *url.URL) {
	return
}

func (u *Exporter) createOutputDir(dir string) error {
	return os.MkdirAll(dir, 0770)
}
