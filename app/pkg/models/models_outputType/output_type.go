package models_outputType

type OutputType string

const (
	Undefined                   OutputType = ""
	Facebook                    OutputType = "facebook"
	WoocommercePluginProductCsv OutputType = "woo-product-csv"
	WoocommercePluginWebToffee  OutputType = "woo-webtoffee"
)

func (e *OutputType) String() string {
	return string(*e)
}

func (e *OutputType) Set(v string) error {
	*e = OutputType(v)
	return nil
}

func (e *OutputType) Type() string {
	return "OutputType"
}

func GetAllSupportedValues() string {
	return "[facebook|woo-product-csv|woo-webtoffee]"
}
