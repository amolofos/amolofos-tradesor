package unloader

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/url"
	"os"
	"reflect"
	"strings"

	"amolofos/tradesor/pkg/models/models_csv"
)

type Unloader struct {
	replacer *strings.Replacer
}

func (u *Unloader) Init() {
	u.replacer = strings.NewReplacer(" ", "", "/", "")
}

func (u *Unloader) Unload(csvDoc *models_csv.Csv, outputTo string) (err error) {
	//remoteUrl, err := url.ParseRequestURI(outputTo)
	//if err != nil {
	//	u.unloadToLocalDir(csvDoc, outputTo)
	//	return
	//}
	//
	//u.unloadToWordpress(csvDoc, remoteUrl)
	u.unloadToLocalDir(csvDoc, outputTo)

	return
}

func (u *Unloader) unloadToLocalDir(csvDoc *models_csv.Csv, outputTo string) {
	fmt.Println("outputTo: " + outputTo)
	fmt.Println("csvDoc: " + csvDoc.Products[0].Id)

	csvMap := map[string][][]string{}
	for i, v := range csvDoc.Products {
		fmt.Println("category: " + string(i) + ", product: " + v.Id)

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
		fmt.Println(category)

		fo, errFileOut := os.Create(outputTo + "/" + category + ".csv")
		if errFileOut != nil {
			log.Fatalln("Error opening file:", errFileOut)
		}

		csvWrite := csv.NewWriter(fo)
		csvWrite.WriteAll(data)
		fo.Close()

		if errCsvWrite := csvWrite.Error(); errCsvWrite != nil {
			log.Fatalln("Error writing csv:", errCsvWrite)
		}
	}
}

func (u *Unloader) unloadToWordpress(csvDoc *models_csv.Csv, remoteUrl *url.URL) {
}
