package models_outputType

type OutputType string

const (
	Undefined   OutputType = ""
	Facebook    OutputType = "facebook"
	Woocommerce OutputType = "woocommerce"
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
	return "[facebook|woocommerce]"
}
