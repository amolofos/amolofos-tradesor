package loader

import (
	"encoding/xml"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"os"

	"amolofos/tradesor/pkg/conf"

	"amolofos/tradesor/pkg/features/tradesor/models"
)

type Loader struct {
}

func (l *Loader) Init() {

}

func (l *Loader) Load(catalog string) (doc *models.Xml, err error) {
	_, errParseUrl := url.ParseRequestURI(catalog)
	if errParseUrl != nil {
		return l.loadFromLocalFile(catalog)
	}

	file, errDownload := l.downloadFileFromUrl(catalog)
	if errDownload != nil {
		return nil, errDownload
	}

	return l.loadFromLocalFile(file)
}

func (l *Loader) downloadFileFromUrl(catalog string) (file string, err error) {
	resp, errGet := http.Get(catalog)
	if errGet != nil {
		return "", errGet
	}
	defer resp.Body.Close()

	f, errFile := os.Create(conf.LOCAL_BUILD_DIR + "/test.xml")
	if errFile != nil {
		return "", errFile
	}
	defer f.Close()

	_, err = io.Copy(f, resp.Body)
	return conf.LOCAL_BUILD_DIR + "/test.xml", nil
}

func (l *Loader) loadFromLocalFile(file string) (doc *models.Xml, err error) {
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

	var xmlDoc models.Xml
	xml.Unmarshal(xmlRead, &xmlDoc.Tradesor)
	return
}
