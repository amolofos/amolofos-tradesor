package models

type Xml struct {
	Tradesor XmlTradesor `xml:"tradesor"`
}

type XmlTradesor struct {
	CreatedAt string      `xml:"created_at"`
	Products  XmlProducts `xml:"products"`
}

type XmlProducts struct {
	Product []XmlProduct `xml:"product"`
}

type XmlProduct struct {
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
