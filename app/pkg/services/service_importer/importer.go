package service_importer

import (
	"encoding/xml"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"github.com/amolofos/tradesor/pkg/conf"
	"github.com/amolofos/tradesor/pkg/features/tradesor/tradesor_models"
)

type Importer struct {
	localFile string
}

func (l *Importer) Init() {
	l.localFile = filepath.Join(conf.LOCAL_BUILD_DIR, conf.LOCAL_FILE)
}

func (l *Importer) Import(catalog string) (xmlDoc *tradesor_models.Xml, err error) {
	slog.Info(fmt.Sprintf("Importing catalog %s.", catalog))

	_, errParseUrl := url.ParseRequestURI(catalog)
	if errParseUrl != nil {
		return l.importFromLocalFile(catalog)
	}

	file, errDownload := l.downloadFileFromUrl(catalog)
	if errDownload != nil {
		return nil, errDownload
	}

	return l.importFromLocalFile(file)
}

func (l *Importer) downloadFileFromUrl(catalog string) (file string, err error) {
	resp, errGet := http.Get(catalog)
	if errGet != nil {
		slog.Error("Error downloading catalog: ", catalog, " with error: ", errGet)
		return "", errGet
	}
	defer resp.Body.Close()

	f, errFile := os.Create(l.localFile)
	if errFile != nil {
		slog.Error("Error storing the catalog locally in ", l.localFile, " with error: ", errFile)
		return "", errFile
	}
	defer f.Close()

	_, err = io.Copy(f, resp.Body)
	return l.localFile, err
}

func (l *Importer) importFromLocalFile(file string) (xmlDoc *tradesor_models.Xml, err error) {
	xmlFile, errXmlOpen := os.Open(file)
	if errXmlOpen != nil {
		slog.Error("Error opening file:", errXmlOpen)
		return nil, errXmlOpen
	}
	defer xmlFile.Close()

	xmlRead, errXmlRead := io.ReadAll(xmlFile)
	if errXmlRead != nil {
		slog.Error("Error reading file:", errXmlRead)
		return nil, errXmlRead
	}

	xmlDoc = &tradesor_models.Xml{}
	xml.Unmarshal(xmlRead, &xmlDoc.Tradesor)

	xmlProducts := xmlDoc.Tradesor.Products.ProductList
	slog.Info(fmt.Sprintf("Imported %d products.", len(xmlProducts)))
	return
}
