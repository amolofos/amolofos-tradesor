package facebook_services

import (
	"fmt"
	"log/slog"
	"math"
	"strings"

	"github.com/amolofos/tradesor/pkg/features/facebook/facebook_conf"
	"github.com/amolofos/tradesor/pkg/features/facebook/facebook_models"
	"github.com/amolofos/tradesor/pkg/features/tradesor/tradesor_models"

	"github.com/amolofos/tradesor/pkg/interfaces/canonical_models"
)

type Facebook struct {
	replacer *strings.Replacer
}

func (f *Facebook) Init() {
	f.replacer = strings.NewReplacer(" ", "", "/", "", ">", "_")
}

func (f *Facebook) Transform(xmlDoc *tradesor_models.Xml) (doc canonical_models.CanonicalModel, err error) {
	categoriesMap := map[string]bool{}
	categories := []string{}

	facebookDoc := &facebook_models.Facebook{}
	facebookDoc.Init()

	xmlProducts := xmlDoc.Tradesor.Products.ProductList

	for _, v := range xmlProducts {
		category := f.replacer.Replace(v.Category)
		_, existsCategory := categoriesMap[category]
		if !existsCategory {
			categoriesMap[category] = true
			categories = append(categories, category)
		}

		availability, existsAvailability := facebook_conf.StockStatusMap[v.StockStatus]
		if !existsAvailability {
			availability = facebook_conf.StockStatusMap["outofstock"]
		}

		titleMaxLength := int(math.Min(float64(len(v.Name)), facebook_conf.TITLE_MAX_LENGTH))
		descriptionMaxLength := int(math.Min(float64(len(v.Name)), facebook_conf.DESCRIPTION_MAX_LENGTH))
		brandMaxLength := int(math.Min(float64(len(v.Manufacturer)), facebook_conf.BRAND_MAX_LENGTH))

		facebookDoc.Products = append(facebookDoc.Products, facebook_models.FacebookProduct{
			Category:            category,
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

	facebookDoc.Categories = categories
	doc = facebookDoc

	slog.Info(fmt.Sprintf("Transformed %d products from %d categories.", len(xmlProducts), len(categories)))

	return
}
