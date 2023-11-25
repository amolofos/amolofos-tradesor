package facebook_services

import (
	"amolofos/tradesor/pkg/features/facebook/facebook_conf"
	"amolofos/tradesor/pkg/features/tradesor/tradesor_models"
	"amolofos/tradesor/pkg/models/models_csv"
	"fmt"
	"log/slog"
	"math"
	"strings"
)

type Facebook struct {
	replacer *strings.Replacer
}

func (f *Facebook) Init() {
	f.replacer = strings.NewReplacer(" ", "", "/", "")
}

func (f *Facebook) TransformToCsv(xmlDoc *tradesor_models.Xml) (csvDoc *models_csv.Csv, err error) {
	categories := map[string]int{}

	csvDoc = &models_csv.Csv{
		Header: facebook_conf.CsvHeader,
	}

	xmlProducts := xmlDoc.Tradesor.Products.ProductList

	for _, v := range xmlProducts {
		category := f.replacer.Replace(v.Category)
		_, existsCategory := categories[category]
		if !existsCategory {
			categories[category] = 0
		} else {
			categories[category] = categories[category] + 1
		}

		availability, existsAvailability := facebook_conf.StockStatusMap[v.StockStatus]
		if !existsAvailability {
			availability = facebook_conf.StockStatusMap["outofstock"]
		}

		titleMaxLength := int(math.Min(float64(len(v.Name)), facebook_conf.TITLE_MAX_LENGTH))
		descriptionMaxLength := int(math.Min(float64(len(v.Name)), facebook_conf.DESCRIPTION_MAX_LENGTH))
		brandMaxLength := int(math.Min(float64(len(v.Manufacturer)), facebook_conf.BRAND_MAX_LENGTH))

		csvDoc.Products = append(csvDoc.Products, models_csv.CsvProduct{
			Id:                  v.Id,
			Title:               v.Name[:titleMaxLength],
			Description:         v.Name[:descriptionMaxLength],
			Availability:        availability,
			Condition:           "new",
			Price:               fmt.Sprintf(facebook_conf.PRICE_FMT_PATTERN, v.SuggestedRetailPrice, "EUR"),
			Link:                "",
			ImageLink:           v.Image,
			Brand:               v.Manufacturer[:brandMaxLength],
			Color:               v.Color,
			ShippingWeight:      v.Weight,
			RichTextDescription: v.Content,
			FbProductCategory:   v.Category,
		})
	}

	slog.Info(fmt.Sprintf("Transformed %d products from %d categories.", len(xmlProducts), len(categories)))
	return
}
