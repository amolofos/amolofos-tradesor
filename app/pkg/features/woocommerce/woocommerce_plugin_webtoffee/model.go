package woocommerce_plugin_webtoffee

import (
	"errors"
	"fmt"
	"strings"

	"github.com/amolofos/tradesor/pkg/models/models_outputFormat"

	"github.com/gocarina/gocsv"
)

type WoocommerceModel struct {
	header     []string
	categories []string

	products []Product
}

/*
 * See the following resources
 *   * https://wordpress.org/plugins/product-import-export-for-woo/#description
 *   * https://www.webtoffee.com/wp-content/uploads/2021/05/Basic-Product_WooCommerce_Sample_CSV.csv
 *
 * E.g.
 *   * parent_sku,sku,post_title,post_excerpt,post_content,post_status,regular_price,sale_price,stock_status,stock,manage_stock,weight,Images,tax:product_type,tax:product_cat,tax:product_tag
 *   * ,A11,Samsung Galaxy Ace Duos,"-5 MP Primary Camera -3.5 inch Capacitive Touchscreen -Wi-Fi Enabled","Having a much vaunted reputation of being one of the world's leading budget smartphone makers,<br/> the Samsung Galaxy Ace Duos S6802 does quite well to maintain Samsung's authority in the current market. The device comes in a bar shape and has a sleek look to it. The Ace Duos S6802 is a Smart Dual Active SIM, (GSM + GSM) device with both the inserted SIM cards being active at the same time. If you are on a call on sim 1 and you get a call on sim 2, the call would be diverted to sim 1 and the display would show it as call waiting.The same will be apply for both the Sims. So that we cannot miss out any calls on sim 1 and sim 2. This Samsung phone is run on the Android v2.3 (Gingerbread) OS with a snappy processor speed of 832 MHz. To assure the smooth running of apps, the device comes with an adequate 512 MB of RAM.",publish,11.99,10.99,instock,10,yes,0.5,,simple,Mobile Phone,Smart Phone|Mobile
 *   * ,A12,Formal shoe,Formal shoe leather made,"Formal shoe leather made , black color",publish,10,8,instock,20,yes,0.5,,simple,Shoes|Men’s Shoes,Foot Wears
 */
type Product struct {
	Category  string `csv:"-"`
	Id        string `csv:"ID"`
	ParentId  string `csv:"Parent"`
	Sku       string `csv:"SKU"`
	ParentSku string `csv:"Parent_SKU"`

	PostTitle    string `csv:"post_title"`
	PostExcerpt  string `csv:"post_excerpt"`
	PostContent  string `csv:"post_content"`
	PostStatus   string `csv:"post_status"`
	RegularPrice string `csv:"regular_price"`
	SalePrice    string `csv:"sale_price"`
	StockStatus  string `csv:"stock_status"`
	Stock        string `csv:"stock"`
	ManageStock  string `csv:"manage_stock"`
	Weight       string `csv:"weight"`
	Images       string `csv:"Images"`

	Visibility       string `csv:"visibility"`
	Backorders       string `csv:"backorders"`
	SoldIndividually string `csv:"sold_individually"`

	Upsells     string `csv:"Upsells"`
	CrossSells  string `csv:"Cross-sells"`
	ExternalUrl string `csv:"External URL"`
	ButtonText  string `csv:"Button text"`
	Position    string `csv:"Position"`

	TypeDownloadable string `csv:"Downloadable"`
	TypeVirtual      string `csv:"Virtual"`

	TaxProductType string `csv:"tax:product_type"`
	TaxProductCat  string `csv:"tax:product_cat"`
	TaxProductTag  string `csv:"tax:product_tag"`

	TaxStatus string `csv:"Tax status"`
	TaxClass  string `csv:"Tax class"`
}

func NewWoocommerceModel() (w *WoocommerceModel) {
	w = &WoocommerceModel{}
	w.Init()
	return
}

func (w *WoocommerceModel) Init() {}

func (w *WoocommerceModel) Header() []string {
	return []string{}
}

func (w *WoocommerceModel) Categories() []string {
	return w.categories
}

func (w *WoocommerceModel) ProductIds(category string) (productIds []string, err error) {
	productIds = []string{}

	for _, v := range w.products {
		if v.Category == category {
			productIds = append(productIds, v.Id)
		}
	}

	return
}

func (w *WoocommerceModel) Products(category string) (products [][]string, err error) {
	products = [][]string{}

	for _, v := range w.products {
		products = append(products, []string{
			v.Id,
		})
	}

	return
}

func (w *WoocommerceModel) productsList(category string) (products []Product, err error) {
	products = []Product{}

	for _, v := range w.products {
		if v.Category == category {
			products = append(products, v)
		}
	}

	return
}

func (w *WoocommerceModel) ExportHeader(outputFormat models_outputFormat.OutputFormat) (header string, err error) {
	headerList := w.Header()

	switch outputFormat {
	case models_outputFormat.CSV:
		header = strings.Trim(strings.Join(strings.Fields(fmt.Sprint(headerList)), ","), "[]")
	default:
		errStr := fmt.Sprintf("WoocommerceModel: Output format %s is not supported.", outputFormat)
		err = errors.New(errStr)
	}

	return
}

func (w *WoocommerceModel) Export(category string, outputFormat models_outputFormat.OutputFormat) (nProducts int, products string, err error) {
	var productsList []Product

	productsList, err = w.productsList(category)
	if err != nil {
		return
	}

	nProducts = len(productsList)

	switch outputFormat {
	case models_outputFormat.CSV:
		products, err = gocsv.MarshalString(&productsList)
	default:
		errStr := fmt.Sprintf("WoocommerceModel: Output format %s is not supported.", outputFormat)
		err = errors.New(errStr)
	}

	return
}
