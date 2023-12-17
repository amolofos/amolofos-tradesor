package woocommerce_plugin_product_csv

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
 *   * https://github.com/woocommerce/woocommerce/wiki/Product-CSV-Import-Schema#csv-columns-and-formatting
 *   * https://github.com/woocommerce/woocommerce/blob/trunk/plugins/woocommerce/sample-data/sample_products.csv
 *
 * E.g.
 *   ID,Type,SKU,Name,Published,"Is featured?","Visibility in catalog","Short description",Description,"Date sale price starts","Date sale price ends","Tax status","Tax class","In stock?",Stock,"Backorders allowed?","Sold individually?","Weight (lbs)","Length (in)","Width (in)","Height (in)","Allow customer reviews?","Purchase note","Sale price","Regular price",Categories,Tags,"Shipping class",Images,"Download limit","Download expiry days",Parent,"Grouped products",Upsells,Cross-sells,"External URL","Button text",Position,"Attribute 1 name","Attribute 1 value(s)","Attribute 1 visible","Attribute 1 global","Attribute 2 name","Attribute 2 value(s)","Attribute 2 visible","Attribute 2 global","Meta: _wpcom_is_markdown","Download 1 name","Download 1 URL","Download 2 name","Download 2 URL"
 *   44,variable,woo-vneck-tee,"V-Neck T-Shirt",1,1,visible,"This is a variable product.","Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. Vestibulum tortor quam, feugiat vitae, ultricies eget, tempor sit amet, ante. Donec eu libero sit amet quam egestas semper. Aenean ultricies mi vitae est. Mauris placerat eleifend leo.",,,taxable,,1,,0,0,.5,24,1,2,1,,,,"Clothing > Tshirts",,,"https://woocommercecore.mystagingwebsite.com/wp-content/uploads/2017/12/vneck-tee-2.jpg, https://woocommercecore.mystagingwebsite.com/wp-content/uploads/2017/12/vnech-tee-green-1.jpg, https://woocommercecore.mystagingwebsite.com/wp-content/uploads/2017/12/vnech-tee-blue-1.jpg",,,,,,,,,0,Color,"Blue, Green, Red",1,1,Size,"Large, Medium, Small",1,1,1,,,,
 */
type Product struct {
	Category             string `csv:"-"`
	Id                   string `csv:"ID"`
	Type                 string `csv:"Type"`
	Sku                  string `csv:"SKU"`
	Name                 string `csv:"Name"`
	Published            string `csv:"Published"`
	IsFeatured           string `csv:"Is featured?"`
	VisibilityInCatalog  string `csv:"Visibility in catalog"`
	ShortDescription     string `csv:"Short description"`
	Description          string `csv:"Description"`
	DateSalePriceStarts  string `csv:"Date sale price starts"`
	DateSalePriceEnds    string `csv:"Date sale price ends"`
	TaxStatus            string `csv:"Tax status"`
	TaxClass             string `csv:"Tax class"`
	InStock              string `csv:"In stock?"`
	Stock                string `csv:"Stock"`
	BackordersAllowed    string `csv:"Backorders allowed?"`
	SoldIndividually     string `csv:"Sold individually?"`
	WeightKg             string `csv:"Weight (kn)"`
	LengthCm             string `csv:"Length (cm)"`
	WidthCm              string `csv:"Width (cm)"`
	HeightCm             string `csv:"Height (cm)"`
	AllowCustomerReviews string `csv:"Allow customer reviews?"`
	PurchaseNote         string `csv:"Purchase note"`
	SalePrice            string `csv:"Sale price"`
	RegularPrice         string `csv:"Regular price"`
	Categories           string `csv:"Categories"`
	Tags                 string `csv:"Tags"`
	ShippingClass        string `csv:"Shipping class"`
	Images               string `csv:"Images"`
	DownloadLimit        string `csv:"Download limit"`
	DownloadExpiryDays   string `csv:"Download expiry days"`
	Parent               string `csv:"Parent"`
	GroupedProducts      string `csv:"Grouped products"`
	Upsells              string `csv:"Upsells"`
	CrossSells           string `csv:"Cross-sells"`
	ExternalUrl          string `csv:"External URL"`
	ButtonText           string `csv:"Button text"`
	Position             string `csv:"Position"`
	Attribute1Name       string `csv:"Attribute 1 name"`
	Attribute1Values     string `csv:"Attribute 1 value(s)"`
	Attribute1Visible    string `csv:"Attribute 1 visible"`
	Attribute1Global     string `csv:"Attribute 1 global"`
	Attribute2Name       string `csv:"Attribute 2 name"`
	Attribute2Values     string `csv:"Attribute 2 value(s)"`
	Attribute2Visible    string `csv:"Attribute 2 visible"`
	Attribute2Global     string `csv:"Attribute 2 global"`
	MetaWwpcomIsMarkdown string `csv:"Meta: _wpcom_is_markdown"`
	Download1Name        string `csv:"Download 1 name"`
	Download1Url         string `csv:"Download 1 URL"`
	Download2Name        string `csv:"Download 2 name"`
	Download2Url         string `csv:"Download 2 URL"`
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
