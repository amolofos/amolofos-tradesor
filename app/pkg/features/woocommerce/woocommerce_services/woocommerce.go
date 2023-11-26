package woocommerce_services

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/amolofos/tradesor/pkg/features/tradesor/tradesor_models"
	"github.com/amolofos/tradesor/pkg/features/woocommerce/woocommerce_models"

	"github.com/amolofos/tradesor/pkg/interfaces/canonical_models"
)

type Woocommerce struct {
	replacer *strings.Replacer
}

func (w *Woocommerce) Init() {
	w.replacer = strings.NewReplacer(" ", "", "/", "", ">", "_")
}

func (w *Woocommerce) Transform(xmlDoc *tradesor_models.Xml) (doc canonical_models.CanonicalModel, err error) {
	categories := map[string]int{}

	woocommerceDoc := &woocommerce_models.Woocommerce{}
	woocommerceDoc.Init()

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
	doc = woocommerceDoc
	return
}
