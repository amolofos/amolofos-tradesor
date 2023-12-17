package woocommerce_plugin_webtoffee

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
			Category:     category,
			Id:           v.Id,
			ParentId:     v.ParentProductID,
			PostTitle:    v.Name,
			PostExcerpt:  "",
			PostContent:  v.Content,
			PostStatus:   "publish",
			RegularPrice: v.RegularPrice,
			SalePrice:    v.SuggestedRetailPrice,
			StockStatus:  v.StockStatus,
			Stock:        v.Stock,
			ManageStock:  "yes",
			Weight:       v.Weight,
			Images:       v.Gallery,

			Visibility:       "1",
			Backorders:       "0",
			SoldIndividually: "0",

			TypeDownloadable: "0",
			TypeVirtual:      "0",

			TaxProductType: "",
			TaxProductCat:  v.Category,
			TaxProductTag:  "",
		})
	}

	woocommerceDoc.categories = categories
	nProducts = len(woocommerceDoc.products)
	doc = woocommerceDoc

	slog.Info(fmt.Sprintf("WoocommerceService: Transformed %d products from %d categories.", nProducts, len(categories)))
	return
}
