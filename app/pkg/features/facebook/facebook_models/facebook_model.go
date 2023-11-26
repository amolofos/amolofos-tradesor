package facebook_models

import (
	"github.com/amolofos/tradesor/pkg/features/facebook/facebook_conf"

	"github.com/amolofos/tradesor/pkg/models/models_outputFormat"

	"github.com/gocarina/gocsv"
)

type Facebook struct {
	Header     []string
	Categories []string

	Products []FacebookProduct
}

type FacebookProduct struct {
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

func (f *Facebook) Init() {
	f.Header = facebook_conf.Header
}

func (f *Facebook) GetHeader() []string {
	return f.Header
}

func (f *Facebook) GetCategories() []string {
	return f.Categories
}

func (f *Facebook) GetProductsFormatted(format models_outputFormat.OutputFormat) (output string, err error) {
	switch format {
	case models_outputFormat.CSV:
		output, err = gocsv.MarshalString(f.Products)
	}

	return
}

func (f *Facebook) GetProductsForCategory(category string) (products [][]string) {
	products = [][]string{}

	for _, v := range f.Products {
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

func (f *Facebook) GetProductsForCategoryFormatted(format models_outputFormat.OutputFormat, category string) (output string, err error) {
	products := []FacebookProduct{}

	for _, v := range f.Products {
		if v.Category == category {
			products = append(products, v)
		}
	}

	switch format {
	case models_outputFormat.CSV:
		output, err = gocsv.MarshalString(products)
	}

	return
}
