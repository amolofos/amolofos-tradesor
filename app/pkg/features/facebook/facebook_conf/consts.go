package facebook_conf

const TITLE_MAX_LENGTH = 65
const DESCRIPTION_MAX_LENGTH = 9999
const BRAND_MAX_LENGTH = 100
const PRICE_FMT_PATTERN = "%s %s"

var Header = []string{
	"id",
	"title",
	"description",
	"availability",
	"condition",
	"price",
	"link",
	"image_link",
	"brand",
	"color",
	"shipping_weight",
	"rich_text_description",
	"fb_product_category",
}

var StockStatusMap = map[string]string{
	"outofstock": "out of stock",
	"instock":    "in stock",
}
