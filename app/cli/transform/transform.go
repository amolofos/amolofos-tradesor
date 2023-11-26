package transform

import (
	"log/slog"

	"github.com/spf13/cobra"

	"github.com/amolofos/tradesor/pkg/models/models_outputFormat"
	"github.com/amolofos/tradesor/pkg/models/models_outputType"
	"github.com/amolofos/tradesor/pkg/services/service_exporter"
	"github.com/amolofos/tradesor/pkg/services/service_importer"
	"github.com/amolofos/tradesor/pkg/services/service_transformer"
)

var (
	outputTo     string = ""
	outputFormat models_outputFormat.OutputFormat
	outputType   models_outputType.OutputType
)

var TransformCmd = &cobra.Command{
	Use:   "transform [--outputFormat xml|csv] [--outputType facebook|woocommerce]",
	Short: "transform Tradesor catalog",
	Long:  "transform Tradesor catalog",
	Run:   transform,
}

func init() {
	TransformCmd.Flags().StringVarP(&outputTo, "outputTo", "", "./output", "Location of the output files. It can be a url (for wordpress) or a local directory.")
	TransformCmd.Flags().VarP(&outputType, "outputType", "", "What type to use for the output: "+models_outputType.GetAllSupportedValues())
	TransformCmd.Flags().VarP(&outputFormat, "outputFormat", "", "What format to use for the output: "+models_outputFormat.GetAllSupportedValues())
}

func transform(cmd *cobra.Command, args []string) {
	var importer = &service_importer.Importer{}
	importer.Init()

	var transformer = &service_transformer.Transformer{}
	transformer.Init()

	var exporter = &service_exporter.Exporter{}
	exporter.Init()

	doc, errImport := importer.Import(cmd.Flag("catalog").Value.String())
	if errImport != nil {
		slog.Error("Failed to import document with error ", errImport)
	}

	out, errTransform := transformer.Transform(doc, models_outputType.OutputType(cmd.Flag("outputType").Value.String()))
	if errTransform != nil {
		slog.Error("Failed to transform document with error ", errTransform)
	}

	errExport := exporter.Export(out, models_outputFormat.OutputFormat(cmd.Flag("outputFormat").Value.String()), cmd.Flag("outputTo").Value.String())
	if errExport != nil {
		slog.Error("Failed to export document with error ", errExport)
	}
}
