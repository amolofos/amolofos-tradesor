package models_outputFormat

type OutputFormat string

const (
	Undefined OutputFormat = ""
	Facebook  OutputFormat = "facebook"
	Wordpress OutputFormat = "wordpress"
)

func (e *OutputFormat) String() string {
	return string(*e)
}

func (e *OutputFormat) Set(v string) error {
	*e = OutputFormat(v)
	return nil
}

func (e *OutputFormat) Type() string {
	return "OutputFormat"
}

func GetAllSupportedValues() string {
	return "[facebook|wordpress]"
}
