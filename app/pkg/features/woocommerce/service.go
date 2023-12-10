package woocommerce

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
	categories := map[string]int{}

	woocommerceDoc := NewWoocommerceModel()

	xmlProducts := xmlDoc.Tradesor.Products.ProductList

	for _, v := range xmlProducts {
		category := w.replacer.Replace(v.Category)
		_, existsCategory := categories[category]
		if !existsCategory {
			categories[category] = 0
		} else {
			categories[category] = categories[category] + 1
		}
	}

	slog.Info(fmt.Sprintf("Transformed %d products from %d categories.", len(xmlProducts), len(categories)))
	nProducts = len(xmlProducts)
	doc = woocommerceDoc
	return
}

// Exporter interface.
func (w *WoocommerceService) Write(doc canonical_models.CanonicalModel) (err error) {
	return
}
