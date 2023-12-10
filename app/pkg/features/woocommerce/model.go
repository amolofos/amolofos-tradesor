// https://github.com/woocommerce/woocommerce/wiki/Product-CSV-Import-Schema#csv-columns-and-formatting
// https://github.com/woocommerce/woocommerce/tree/trunk/plugins/woocommerce/sample-data
package woocommerce

import "github.com/amolofos/tradesor/pkg/models/models_outputFormat"

type WoocommerceModel struct {
	header     []string
	categories []string

	products []Product
}

type Product struct {
	Category string `csv:"-"`
	Id       string `csv:"id"`
}

func NewWoocommerceModel() (w *WoocommerceModel) {
	w = &WoocommerceModel{}
	w.Init()
	return
}

func (w *WoocommerceModel) Init() {
	w.header = Header
}

func (w *WoocommerceModel) Header() []string {
	return w.header
}

func (w *WoocommerceModel) Categories() []string {
	return w.categories
}

func (w *WoocommerceModel) ProductIds(category string) (productIds []string) {
	productIds = []string{}

	for _, v := range w.products {
		if v.Category == category {
			productIds = append(productIds, v.Id)
		}
	}

	return productIds
}

func (w *WoocommerceModel) Products(category string) (products [][]string) {
	products = [][]string{}

	for _, v := range w.products {
		products = append(products, []string{
			v.Id,
		})
	}

	return products
}

func (w *WoocommerceModel) FormatProduct(productId string, format models_outputFormat.OutputFormat) (output string, err error) {
	return
}

func (w *WoocommerceModel) FormatProducts(category string, format models_outputFormat.OutputFormat) (output string, err error) {
	return
}
