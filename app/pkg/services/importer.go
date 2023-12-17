package services

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
	"github.com/amolofos/tradesor/pkg/features/tradesor"
)

type Importer struct {
	localFile string
}

func NewImporter() (i *Importer) {
	i = &Importer{}
	i.localFile = filepath.Join(conf.LOCAL_BUILD_DIR, conf.LOCAL_FILE)
	return
}

func (i *Importer) Import(catalog string) (nProducts int, xmlDoc *tradesor.ModelXml, err error) {
	slog.Info(fmt.Sprintf("Importer: Importing catalog %s.", catalog))

	_, errParseUrl := url.ParseRequestURI(catalog)
	if errParseUrl != nil {
		return i.importFromLocalFile(catalog)
	}

	file, errDownload := i.downloadFileFromUrl(catalog)
	if errDownload != nil {
		err = errDownload
		return
	}

	return i.importFromLocalFile(file)
}

func (i *Importer) downloadFileFromUrl(catalog string) (file string, err error) {
	resp, errGet := http.Get(catalog)
	if errGet != nil {
		slog.Error("Importer: Error downloading catalog: ", catalog, " with error: ", errGet)
		err = errGet

		return
	}
	defer resp.Body.Close()

	errDir := os.MkdirAll(filepath.Dir(i.localFile), 0700)
	if errDir != nil {
		slog.Error("Importer: Error creating directory for file ", i.localFile, " with error: ", errDir)
		err = errDir

		return
	}

	f, errFile := os.Create(i.localFile)
	if errFile != nil {
		slog.Error("Importer: Error creating the file catalog locally as ", i.localFile, " with error: ", errFile)
		err = errFile

		return
	}
	defer f.Close()

	_, err = io.Copy(f, resp.Body)
	return i.localFile, err
}

func (i *Importer) importFromLocalFile(file string) (nProducts int, xmlDoc *tradesor.ModelXml, err error) {
	xmlFile, errXmlOpen := os.Open(file)
	if errXmlOpen != nil {
		slog.Error("Importer: Error opening file:", errXmlOpen)
		err = errXmlOpen
		return
	}
	defer xmlFile.Close()

	xmlRead, errXmlRead := io.ReadAll(xmlFile)
	if errXmlRead != nil {
		slog.Error("Importer: Error reading file:", errXmlRead)
		err = errXmlRead
		return
	}

	xmlDoc = &tradesor.ModelXml{}
	xml.Unmarshal(xmlRead, &xmlDoc.Tradesor)

	xmlProducts := xmlDoc.Tradesor.Products.ProductList
	nProducts = len(xmlProducts)
	slog.Info(fmt.Sprintf("Importer: Imported %d products.", nProducts))
	return
}
