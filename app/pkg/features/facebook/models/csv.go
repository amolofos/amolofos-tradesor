package models

type Csv struct {
	Products []CsvProduct
}

type CsvProduct struct {
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
