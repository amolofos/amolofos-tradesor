// https://github.com/woocommerce/woocommerce/wiki/Product-CSV-Import-Schema#csv-columns-and-formatting
// https://github.com/woocommerce/woocommerce/tree/trunk/plugins/woocommerce/sample-data
package woocommerce_models

import (
	"github.com/amolofos/tradesor/pkg/features/woocommerce/woocommerce_conf"
	"github.com/amolofos/tradesor/pkg/models/models_outputFormat"
)

type Woocommerce struct {
	Header []string

	Products []WoocommerceProduct
}

type WoocommerceProduct struct {
	Id string `csv:"id"`
}

func (w *Woocommerce) Init() {
	w.Header = woocommerce_conf.Header
}

func (w *Woocommerce) GetHeader() []string {
	return w.Header
}

func (w *Woocommerce) GetCategories() []string {
	return []string{}
}

func (w *Woocommerce) GetProductsFormatted(format models_outputFormat.OutputFormat) (output string, err error) {
	return "", nil
}

func (w *Woocommerce) GetProductsForCategory(category string) [][]string {
	return [][]string{}
}

func (w *Woocommerce) GetProductsForCategoryFormatted(format models_outputFormat.OutputFormat, category string) (output string, err error) {
	return "", nil
}
