package facebook

import (
	"errors"
	"fmt"
	"strings"

	"github.com/amolofos/tradesor/pkg/models/models_outputFormat"
	"github.com/gocarina/gocsv"
)

type FacebookModel struct {
	header     []string
	categories []string

	products []Product
}

type Product struct {
	Category            string `csv:"-"`
	Id                  string `csv:"id"`
	Title               string `csv:"title"`
	Description         string `csv:"description"`
	Availability        string `csv:"availability"`
	Condition           string `csv:"condition"`
	Price               string `csv:"price"`
	Link                string `csv:"link"`
	ImageLink           string `csv:"image_link"`
	Brand               string `csv:"brand"`
	Color               string `csv:"color"`
	ShippingWeight      string `csv:"shipping_weight"`
	RichTextDescription string `csv:"rich_text_description"`
	FbProductCategory   string `csv:"fb_product_category"`
}

func NewFacebookModel() (f *FacebookModel) {
	f = &FacebookModel{}
	f.Init()
	return
}

func (f *FacebookModel) Init() {
	f.header = Header
}

func (f *FacebookModel) Header() []string {
	return f.header
}

func (f *FacebookModel) SetHeader(header []string) {
	f.header = header
}

func (f *FacebookModel) Categories() []string {
	return f.categories
}

func (f *FacebookModel) SetCategories(categories []string) {
	f.categories = categories
}

func (f *FacebookModel) ProductIds(category string) (productIds []string, err error) {
	productIds = []string{}

	for _, v := range f.products {
		if v.Category == category {
			productIds = append(productIds, v.Id)
		}
	}

	return
}

func (f *FacebookModel) Products(category string) (products [][]string, err error) {
	products = [][]string{}

	for _, v := range f.products {
		if v.Category == category {
			products = append(products, []string{
				v.Id,
				v.Title,
				v.Description,
				v.Availability,
				v.Condition,
				v.Price,
				v.Link,
				v.ImageLink,
				v.Brand,
				v.Color,
				v.ShippingWeight,
				v.RichTextDescription,
				v.FbProductCategory,
			})
		}
	}

	return
}

func (f *FacebookModel) productsList(category string) (products []Product, err error) {
	products = []Product{}

	for _, v := range f.products {
		if v.Category == category {
			products = append(products, v)
		}
	}

	return
}

func (f *FacebookModel) ExportHeader(outputFormat models_outputFormat.OutputFormat) (header string, err error) {
	headerList := f.Header()

	switch outputFormat {
	case models_outputFormat.CSV:
		header = strings.Trim(strings.Join(strings.Fields(fmt.Sprint(headerList)), ","), "[]")
	default:
		errStr := fmt.Sprintf("FacebookModel: Output format %s is not supported.", outputFormat)
		err = errors.New(errStr)
	}

	return
}

func (f *FacebookModel) Export(category string, outputFormat models_outputFormat.OutputFormat) (nProducts int, products string, err error) {
	var productsList []Product

	productsList, err = f.productsList(category)
	if err != nil {
		return
	}

	nProducts = len(productsList)

	switch outputFormat {
	case models_outputFormat.CSV:
		products, err = gocsv.MarshalString(&productsList)
	default:
		errStr := fmt.Sprintf("FacebookModel: Output format %s is not supported.", outputFormat)
		err = errors.New(errStr)
	}

	return
}
