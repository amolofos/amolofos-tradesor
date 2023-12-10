package tradesor

type ModelXml struct {
	Tradesor Xml `xml:"tradesor"`
}

type Xml struct {
	CreatedAt string   `xml:"created_at"`
	Products  Products `xml:"products"`
}

type Products struct {
	ProductList []Product `xml:"product"`
}

type Product struct {
	Id                   string `xml:"id"`
	ParentProductID      string `xml:"ParentProductID"`
	Name                 string `xml:"name"`
	Sku                  string `xml:"Sku"`
	Content              string `xml:"Content"`
	Image                string `xml:"image"`
	Gallery              string `xml:"gallery"`
	Category             string `xml:"category"`
	Marketing            string `xml:"marketing"`
	Mpn                  string `xml:"mpn"`
	RegularPrice         string `xml:"RegularPrice"`
	SuggestedRetailPrice string `xml:"SuggestedRetailPrice"`
	Color                string `xml:"color"`
	Manufacturer         string `xml:"manufacturer"`
	StockStatus          string `xml:"StockStatus"`
	Stock                string `xml:"Stock"`
	Weight               string `xml:"weight"`
	ShippingLeadTime     string `xml:"ShippingLeadTime"`
}

func (p *Product) String() (str string) {
	return p.Category
}
