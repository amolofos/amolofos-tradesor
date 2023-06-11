package main

import (
	"encoding/csv"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"reflect"
	"strings"
)

type Xml struct {
	Tradesor XmlTradesor `xml:"tradesor"`
}

type XmlTradesor struct {
	CreatedAt string      `xml:"created_at"`
	Products  XmlProducts `xml:"products"`
}

type XmlProducts struct {
	Product []XmlProduct `xml:"product"`
}

type XmlProduct struct {
	Id                   string `xml:"id"`
	ParentProductID      string `xml:"ParentProductID"`
	Name                 string `xml:"name"`
	Sku                  string `xml:"Sku"`
	Content              string `xml:"Content"`
	Image                string `xml:"image"`
	Gallery              string `xml:"gallery"`
	Category             string `xml:"category"`
	Marketing            string `xml:"marketing"`
	Mpn                  string `xml:"mpn"`
	RegularPrice         string `xml:"RegularPrice"`
	SuggestedRetailPrice string `xml:"SuggestedRetailPrice"`
	Color                string `xml:"color"`
	Manufacturer         string `xml:"manufacturer"`
	StockStatus          string `xml:"StockStatus"`
	Stock                string `xml:"Stock"`
	Weight               string `xml:"weight"`
	ShippingLeadTime     string `xml:"ShippingLeadTime"`
}

type Csv struct {
	Products []CsvProduct
}

type CsvProduct struct {
	Id                  string `csv:"id"`
	Title               string `csv:"title"`
	Description         string `csv:"description"`
	Availability        string `csv:"availability"`
	Condition           string `csv:"condition"`
	Price               string `csv:"price"`
	Link                string `csv:"link"`
	ImageLink           string `csv:"image_link"`
	Brand               string `csv:"brand"`
	Color               string `csv:"color"`
	ShippingWeight      string `csv:"shipping_weight"`
	RichTextDescription string `csv:"rich_text_description"`
	FbProductCategory   string `csv:"fb_product_category"`
}

const TITLE_MAX_LENGTH = 65
const DESCRIPTION_MAX_LENGTH = 9999
const BRAND_MAX_LENGTH = 100
const PRICE_FMT_PATTERN = "%s %s"

var csvHeader = []string{
	"id",
	"title",
	"description",
	"availability",
	"condition",
	"price",
	"link",
	"image_link",
	"brand",
	"color",
	"shipping_weight",
	"rich_text_description",
	"fb_product_category",
}

var stockStatusMap = map[string]string{
	"outofstock": "out of stock",
	"instock":    "in stock",
}

var replacer = strings.NewReplacer(" ", "", "/", "")

func main() {
	//fileName := "www.tradesor.gr.2023-06.simple"
	fileName := "www.tradesor.gr.2023-06"
	xmlFile, errXmlOpen := os.Open("data/" + fileName + ".xml")
	if errXmlOpen != nil {
		log.Fatalln("Error opening file:", errXmlOpen)
	}
	defer xmlFile.Close()

	xmlRead, errXmlRead := ioutil.ReadAll(xmlFile)
	if errXmlRead != nil {
		log.Fatalln("Error reading file:", errXmlRead)
	}

	var xmlDoc Xml
	var csvDoc Csv
	csvDoc.Products = []CsvProduct{}

	xml.Unmarshal(xmlRead, &xmlDoc.Tradesor)
	//fmt.Println(xmlDoc.Tradesor)

	categories := map[string]int{}

	xmlProducts := xmlDoc.Tradesor.Products.Product
	for _, v := range xmlProducts {
		category := replacer.Replace(v.Category)
		_, existsCategory := categories[category]
		if !existsCategory {
			categories[category] = 0
		} else {
			categories[category] = categories[category] + 1
		}

		availability, existsAvailability := stockStatusMap[v.StockStatus]
		if !existsAvailability {
			availability = stockStatusMap["outofstock"]
		}

		titleMaxLength := int(math.Min(float64(len(v.Name)), TITLE_MAX_LENGTH))
		descriptionMaxLength := int(math.Min(float64(len(v.Name)), DESCRIPTION_MAX_LENGTH))
		brandMaxLength := int(math.Min(float64(len(v.Manufacturer)), BRAND_MAX_LENGTH))

		csvDoc.Products = append(csvDoc.Products, CsvProduct{
			Id:                  v.Id,
			Title:               v.Name[:titleMaxLength],
			Description:         v.Name[:descriptionMaxLength],
			Availability:        availability,
			Condition:           "new",
			Price:               fmt.Sprintf(PRICE_FMT_PATTERN, v.SuggestedRetailPrice, "EUR"),
			Link:                "",
			ImageLink:           v.Image,
			Brand:               v.Manufacturer[:brandMaxLength],
			Color:               v.Color,
			ShippingWeight:      v.Weight,
			RichTextDescription: v.Content,
			FbProductCategory:   v.Category,
		})
	}

	csvMap := map[string][][]string{}
	for _, v := range csvDoc.Products {

		ref := reflect.ValueOf(v)
		row := make([]string, ref.NumField())

		for j := 0; j < ref.NumField(); j++ {
			row[j] = ref.Field(j).String()
		}

		category := replacer.Replace(v.FbProductCategory)

		_, existsCategoryCsv := csvMap[category]
		if !existsCategoryCsv {
			csvMap[category] = [][]string{}
			csvMap[category] = append(csvMap[category], csvHeader)
		}

		csvMap[category] = append(csvMap[category], row)
	}

	for i, v := range csvMap {
		fmt.Println(i)

		fo, errFileOut := os.Create("data/" + fileName + "." + i + ".csv")
		if errFileOut != nil {
			log.Fatalln("Error opening file:", errFileOut)
		}

		csvWrite := csv.NewWriter(fo)
		csvWrite.WriteAll(v)
		fo.Close()

		if errCsvWrite := csvWrite.Error(); errCsvWrite != nil {
			log.Fatalln("Error writing csv:", errCsvWrite)
		}
	}
}
