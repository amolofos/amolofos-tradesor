package transformer

import (
	"amolofos/tradesor/pkg/features/tradesor/models"
	modelFormat "amolofos/tradesor/pkg/models/outputFormat"
	"strings"
)

type Transformer struct {
	replacer *strings.Replacer
}

func (t *Transformer) Init() {
	t.replacer = strings.NewReplacer(" ", "", "/", "")
}

func (t *Transformer) Transform(doc *models.Xml, outputFormat modelFormat.OutputFormat) (out string, err error) {
	switch outputFormat {
	case modelFormat.Facebook:
		t.facebook(doc)
	case modelFormat.Wordpress:
		t.wordpress()
	}

	return "", nil
}

func (t *Transformer) facebook(doc *models.Xml) {

	// categories := map[string]int{}
	//
	// xmlProducts := xmlDoc.Tradesor.Products.Product
	//
	//	for _, v := range xmlProducts {
	//		category := replacer.Replace(v.Category)
	//		_, existsCategory := categories[category]
	//		if !existsCategory {
	//			categories[category] = 0
	//		} else {
	//			categories[category] = categories[category] + 1
	//		}
	//
	//		availability, existsAvailability := stockStatusMap[v.StockStatus]
	//		if !existsAvailability {
	//			availability = stockStatusMap["outofstock"]
	//		}
	//
	//		titleMaxLength := int(math.Min(float64(len(v.Name)), TITLE_MAX_LENGTH))
	//		descriptionMaxLength := int(math.Min(float64(len(v.Name)), DESCRIPTION_MAX_LENGTH))
	//		brandMaxLength := int(math.Min(float64(len(v.Manufacturer)), BRAND_MAX_LENGTH))
	//
	//		csvDoc.Products = append(csvDoc.Products, CsvProduct{
	//			Id:                  v.Id,
	//			Title:               v.Name[:titleMaxLength],
	//			Description:         v.Name[:descriptionMaxLength],
	//			Availability:        availability,
	//			Condition:           "new",
	//			Price:               fmt.Sprintf(PRICE_FMT_PATTERN, v.SuggestedRetailPrice, "EUR"),
	//			Link:                "",
	//			ImageLink:           v.Image,
	//			Brand:               v.Manufacturer[:brandMaxLength],
	//			Color:               v.Color,
	//			ShippingWeight:      v.Weight,
	//			RichTextDescription: v.Content,
	//			FbProductCategory:   v.Category,
	//		})
	//	}
	//
	// csvMap := map[string][][]string{}
	// for _, v := range csvDoc.Products {
	//
	//		ref := reflect.ValueOf(v)
	//		row := make([]string, ref.NumField())
	//
	//		for j := 0; j < ref.NumField(); j++ {
	//			row[j] = ref.Field(j).String()
	//		}
	//
	//		category := replacer.Replace(v.FbProductCategory)
	//
	//		_, existsCategoryCsv := csvMap[category]
	//		if !existsCategoryCsv {
	//			csvMap[category] = [][]string{}
	//			csvMap[category] = append(csvMap[category], csvHeader)
	//		}
	//
	//		csvMap[category] = append(csvMap[category], row)
	//	}
}

func (t *Transformer) wordpress() {
}
