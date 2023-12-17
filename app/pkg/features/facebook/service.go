package facebook

import (
	"fmt"
	"log/slog"
	"math"
	"strings"

	"github.com/amolofos/tradesor/pkg/features/tradesor"
	"github.com/amolofos/tradesor/pkg/interfaces/canonical_models"
)

type FacebookService struct {
	replacer *strings.Replacer
}

func NewFacebookService() (f *FacebookService, err error) {
	f = &FacebookService{}
	f.Init()
	return
}

func (f *FacebookService) Init() {
	f.replacer = strings.NewReplacer(" ", "", "/", "", ">", "_")
}

// Transformer interface.
func (f *FacebookService) CanonicalModel(xmlDoc *tradesor.ModelXml) (nProducts int, doc canonical_models.CanonicalModel, err error) {
	categoriesMap := map[string]bool{}
	categories := []string{}

	facebookDoc := NewFacebookModel()

	xmlProducts := xmlDoc.Tradesor.Products.ProductList

	for _, v := range xmlProducts {
		category := f.replacer.Replace(v.Category)
		_, existsCategory := categoriesMap[category]
		if !existsCategory {
			categoriesMap[category] = true
			categories = append(categories, category)
		}

		availability, existsAvailability := StockStatusMap[v.StockStatus]
		if !existsAvailability {
			availability = StockStatusMap["outofstock"]
		}

		titleMaxLength := int(math.Min(float64(len(v.Name)), TITLE_MAX_LENGTH))
		descriptionMaxLength := int(math.Min(float64(len(v.Name)), DESCRIPTION_MAX_LENGTH))
		brandMaxLength := int(math.Min(float64(len(v.Manufacturer)), BRAND_MAX_LENGTH))

		facebookDoc.products = append(facebookDoc.products, Product{
			Category:            category,
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

	facebookDoc.categories = categories
	nProducts = len(facebookDoc.products)
	doc = facebookDoc

	slog.Info(fmt.Sprintf("FacebookService: Transformed %d products from %d categories.", nProducts, len(categories)))
	return
}
