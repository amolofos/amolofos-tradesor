package woocommerce_plugin_product_csv

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/amolofos/tradesor/pkg/features/tradesor"
	"github.com/amolofos/tradesor/pkg/interfaces/canonical_models"
)

type WoocommerceService struct {
	replacer *strings.Replacer
}

func NewWoocommerceService() (w *WoocommerceService, err error) {
	w = &WoocommerceService{}
	w.Init()
	return
}

func (w *WoocommerceService) Init() {
	w.replacer = strings.NewReplacer(" ", "", "/", "", ">", "_")
}

// Transformer interface.
func (w *WoocommerceService) CanonicalModel(xmlDoc *tradesor.ModelXml) (nProducts int, doc canonical_models.CanonicalModel, err error) {
	categoriesMap := map[string]bool{}
	categories := []string{}

	woocommerceDoc := NewWoocommerceModel()

	xmlProducts := xmlDoc.Tradesor.Products.ProductList

	for _, v := range xmlProducts {
		category := w.replacer.Replace(v.Category)
		_, existsCategory := categoriesMap[category]
		if !existsCategory {
			categoriesMap[category] = true
			categories = append(categories, category)
		}

		woocommerceDoc.products = append(woocommerceDoc.products, Product{
			Category:             category,
			Id:                   v.Id,
			Type:                 "simple",
			Sku:                  v.Sku,
			Name:                 v.Name,
			Published:            "1",
			IsFeatured:           "0",
			VisibilityInCatalog:  "visible",
			ShortDescription:     v.Name,
			Description:          v.Content,
			DateSalePriceStarts:  "",
			DateSalePriceEnds:    "",
			TaxStatus:            "taxable",
			TaxClass:             "",
			InStock:              "1",
			Stock:                v.Stock,
			BackordersAllowed:    "0",
			SoldIndividually:     "0",
			WeightKg:             v.Weight,
			LengthCm:             "",
			WidthCm:              "",
			HeightCm:             "",
			AllowCustomerReviews: "1",
			PurchaseNote:         "",
			SalePrice:            v.SuggestedRetailPrice,
			RegularPrice:         v.RegularPrice,
			Categories:           v.Category,
			Tags:                 "",
			ShippingClass:        "",
			Images:               v.Gallery,
			DownloadLimit:        "",
			DownloadExpiryDays:   "",
			Parent:               "",
			GroupedProducts:      "",
			Upsells:              "",
			CrossSells:           "",
			ExternalUrl:          "",
			ButtonText:           "",
			Position:             "0",
			Attribute1Name:       "",
			Attribute1Values:     "",
			Attribute1Visible:    "",
			Attribute1Global:     "",
			Attribute2Name:       "",
			Attribute2Values:     "",
			Attribute2Visible:    "",
			Attribute2Global:     "",
			MetaWwpcomIsMarkdown: "0",
			Download1Name:        "",
			Download1Url:         "",
			Download2Name:        "",
			Download2Url:         "",
		})
	}

	woocommerceDoc.categories = categories
	nProducts = len(woocommerceDoc.products)
	doc = woocommerceDoc

	slog.Info(fmt.Sprintf("WoocommerceService: Transformed %d products from %d categories.", nProducts, len(categories)))
	return
}
