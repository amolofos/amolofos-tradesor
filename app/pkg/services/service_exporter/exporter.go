package service_exporter

import (
	"encoding/csv"
	"fmt"
	"log/slog"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/amolofos/tradesor/pkg/models/models_csv"
)

type Exporter struct {
	replacer *strings.Replacer
}

func (u *Exporter) Init() {
	u.replacer = strings.NewReplacer(" ", "", "/", "", ">", "_")
}

func (u *Exporter) Export(csvDoc *models_csv.Csv, outputTo string) (err error) {
	//remoteUrl, err := url.ParseRequestURI(outputTo)
	//if err != nil {
	//	u.exportToLocalDir(csvDoc, outputTo)
	//	return
	//}
	//
	//u.exportToWordpress(csvDoc, remoteUrl)
	u.createOutputDir(outputTo)
	u.exportToLocalDir(csvDoc, outputTo)

	return
}

func (u *Exporter) exportToLocalDir(csvDoc *models_csv.Csv, outputTo string) {
	slog.Info(fmt.Sprintf("Exporting %d products.", len(csvDoc.Products)))

	csvMap := map[string][][]string{}
	for _, v := range csvDoc.Products {
		ref := reflect.ValueOf(v)
		row := make([]string, ref.NumField())

		for j := 0; j < ref.NumField(); j++ {
			row[j] = ref.Field(j).String()
		}

		category := u.replacer.Replace(v.FbProductCategory)

		_, existsCategoryCsv := csvMap[category]
		if !existsCategoryCsv {
			csvMap[category] = [][]string{}
			csvMap[category] = append(csvMap[category], csvDoc.Header)
		}

		csvMap[category] = append(csvMap[category], row)
	}

	for category, data := range csvMap {
		categoryFile := filepath.Join(outputTo, category+".csv")

		fo, errFileOut := os.Create(categoryFile)
		if errFileOut != nil {
			slog.Error("Error creating category file: ", categoryFile, " with error: ", errFileOut)
		}

		csvWrite := csv.NewWriter(fo)
		csvWrite.WriteAll(data)
		fo.Close()

		errCsvWrite := csvWrite.Error()
		if errCsvWrite != nil {
			slog.Error("Error writing category file: ", categoryFile, " with error: ", errCsvWrite)
		}
	}

	slog.Info(fmt.Sprintf("Created %d csv files.", len(csvMap)))
}

func (u *Exporter) exportToWordpress(csvDoc *models_csv.Csv, remoteUrl *url.URL) {
	return
}

func (u *Exporter) createOutputDir(dir string) error {
	return os.MkdirAll(dir, 0770)
}
