package transform

import (
	"log/slog"

	"github.com/spf13/cobra"

	"amolofos/tradesor/pkg/models/models_outputFormat"
	"amolofos/tradesor/pkg/services/service_exporter"
	"amolofos/tradesor/pkg/services/service_importer"
	"amolofos/tradesor/pkg/services/service_transformer"
)

var (
	outputTo     string = ""
	outputFormat models_outputFormat.OutputFormat
)

var TransformCmd = &cobra.Command{
	Use:   "transform [--outputFormat facebook]",
	Short: "transform Tradesor catalog",
	Long:  "transform Tradesor catalog",
	Run:   transform,
}

func init() {
	TransformCmd.Flags().StringVarP(&outputTo, "outputTo", "", "./output", "Location of the output files. It can be a url (for wordpress) or a local directory.")
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

	out, errTransform := transformer.Transform(doc, models_outputFormat.OutputFormat(cmd.Flag("outputFormat").Value.String()))
	if errTransform != nil {
		slog.Error("Failed to transform document with error ", errTransform)
	}

	errExport := exporter.Export(out, cmd.Flag("outputTo").Value.String())
	if errExport != nil {
		slog.Error("Failed to export document with error ", errExport)
	}
}
