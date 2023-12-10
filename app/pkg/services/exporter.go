package services

import (
	"errors"
	"fmt"
	"log/slog"
	"net/url"
	"os"
	"strings"

	"github.com/amolofos/tradesor/pkg/interfaces/canonical_models"
	"github.com/amolofos/tradesor/pkg/interfaces/exporter_interfaces"
	"github.com/amolofos/tradesor/pkg/models/models_outputFormat"

	"github.com/amolofos/tradesor/pkg/exporters"
)

type Exporter struct {
	replacer *strings.Replacer
}

func NewExporter() (e *Exporter) {
	e = &Exporter{}
	e.replacer = strings.NewReplacer(" ", "", "/", "", ">", "_")
	return
}

func (e *Exporter) Export(doc canonical_models.CanonicalModel, outputFormat models_outputFormat.OutputFormat, outputTo string) (err error) {
	var exporter exporter_interfaces.Exporter

	exporter, err = e.exporter(outputTo)
	if err != nil {
		slog.Error(fmt.Sprintf("Failed to get exporter for output destination %s with %s error.", outputTo, err.Error()))
		return
	}

	if exporter == nil {
		errStr := fmt.Sprintf("Something went wrong. We could not create an exporter for destination %s.", outputTo)
		slog.Error(errStr)
		err = errors.New(errStr)
		return
	}

	var nProducts int
	nProducts, err = exporter.Write(doc, outputFormat)
	slog.Info(fmt.Sprintf("Exported %d products.", nProducts))
	return
}

func (e *Exporter) exporter(outputTo string) (exporter exporter_interfaces.Exporter, err error) {

	// 1. Do we have a url?
	_, errUri := url.ParseRequestURI(outputTo)
	if errUri == nil {
		slog.Debug("Got a url, creating an http exporter.")
		exporter, err = exporters.NewHttpExporter(outputTo)
		return
	}

	// 2. Or do we have a local directory?
	err = e.createOutputDir(outputTo)
	if err != nil {
		return
	}
	slog.Debug("Got a directory, creating a file exporter.")
	exporter, err = exporters.NewFileExporter(outputTo)

	return
}

func (e *Exporter) createOutputDir(dir string) error {
	return os.MkdirAll(dir, 0770)
}

//
//func (e *Exporter) exportToLocalDir(doc canonical_models.CanonicalModel, outputFormat models_outputFormat.OutputFormat, outputTo string) (err error) {
//
//	categories := doc.Categories()
//	for _, category := range categories {
//		categoryFile := filepath.Join(outputTo, category+".csv")
//		header := doc.Header()
//
//		data, errProducts := doc.FormatProducts(category, outputFormat)
//		if errProducts != nil {
//			slog.Error("Error getting products for category: ", category, " with error: ", errProducts)
//		}
//
//		fo, errFileOut := os.Create(categoryFile)
//		if errFileOut != nil {
//			slog.Error("Error creating category file: ", categoryFile, " with error: ", errFileOut)
//		}
//		defer fo.Close()
//
//		csvWrite := csv.NewWriter(fo)
//		csvWrite.Write(header)
//		errCsvWrite := csvWrite.Error()
//		if errCsvWrite != nil {
//			slog.Error("Error writing category file: ", categoryFile, " with error: ", errCsvWrite)
//		}
//
//		fo.WriteString(data)
//	}
//
//	slog.Info(fmt.Sprintf("Created %d csv files.", len(categories)))
//	return
//}
//
//func (e *Exporter) exportToRemote(doc canonical_models.CanonicalModel, remoteUrl *url.URL) (err error) {
//	return
//}
//
