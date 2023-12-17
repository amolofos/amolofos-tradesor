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
		slog.Error(fmt.Sprintf("Exporter: Failed to get exporter for output destination %s with %s error.", outputTo, err.Error()))
		return
	}

	var nProducts int
	nProducts, err = exporter.Write(doc, outputFormat)
	slog.Info(fmt.Sprintf("Exporter: Exported %d products.", nProducts))
	return
}

func (e *Exporter) exporter(outputTo string) (exporter exporter_interfaces.Exporter, err error) {

	// 1. Do we have a url?
	_, errUri := url.ParseRequestURI(outputTo)
	if errUri == nil {
		slog.Debug("Exporter: Got a url, creating an http exporter.")

		errStr := "Exporter: unfortunately http exporter is not supported yet"
		slog.Error(errStr)
		err = errors.New(errStr)
		return
	}

	// 2. Or do we have a local directory?
	err = e.createOutputDir(outputTo)
	if err != nil {
		return
	}
	slog.Debug("Exporter: Got a directory, creating a file exporter.")
	exporter, err = exporters.NewFileExporter(outputTo)

	return
}

func (e *Exporter) createOutputDir(dir string) error {
	return os.MkdirAll(dir, 0770)
}
