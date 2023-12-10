package facebook

import "github.com/amolofos/tradesor/pkg/models/models_outputFormat"

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

func (f *FacebookModel) Categories() []string {
	return f.categories
}

func (f *FacebookModel) ProductIds(category string) (productIds []string) {
	productIds = []string{}

	for _, v := range f.products {
		if v.Category == category {
			productIds = append(productIds, v.Id)
		}
	}

	return productIds
}

func (f *FacebookModel) Products(category string) (products [][]string) {
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

	return products
}

func (f *FacebookModel) FormatProduct(productId string, format models_outputFormat.OutputFormat) (output string, err error) {
	return
}

func (f *FacebookModel) FormatProducts(category string, format models_outputFormat.OutputFormat) (output string, err error) {
	return
}
