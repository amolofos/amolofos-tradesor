package models_outputFormat

type OutputFormat string

const (
	Undefined OutputFormat = ""
	CSV       OutputFormat = "csv"
)

func (o *OutputFormat) String() string {
	return string(*o)
}

func (o *OutputFormat) Set(v string) error {
	*o = OutputFormat(v)
	return nil
}

func (o *OutputFormat) Type() string {
	return "OutputFormat"
}

func GetAllSupportedValues() string {
	return "[csv]"
}
