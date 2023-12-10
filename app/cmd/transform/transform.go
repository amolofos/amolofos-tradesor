package transform

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/spf13/cobra"

	"github.com/amolofos/tradesor/pkg/models/models_outputFormat"
	"github.com/amolofos/tradesor/pkg/models/models_outputType"
	"github.com/amolofos/tradesor/pkg/services"
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
	var importer = services.NewImporter()
	var transformer = services.NewTransformer()
	var exporter = services.NewExporter()

	nProductsImport, docXml, errImport := importer.Import(cmd.Flag("catalog").Value.String())
	if errImport != nil {
		slog.Error(fmt.Sprintf("Failed to import document with error %s", errImport))
		os.Exit(1)
	}

	if nProductsImport == 0 || docXml == nil {
		slog.Warn("Imported 0 products. Stop processing.")
		return
	}

	nProductsTransform, docCanonical, errTransform := transformer.Transform(
		docXml,
		models_outputType.OutputType(cmd.Flag("outputType").Value.String()),
	)
	if errTransform != nil {
		slog.Error(fmt.Sprintf("Failed to transform document with error %s", errTransform))
		os.Exit(1)
	}

	if nProductsTransform == 0 || docCanonical == nil {
		slog.Warn("Transformed 0 products. Stop processing.")
		return
	}

	errExport := exporter.Export(
		docCanonical,
		models_outputFormat.OutputFormat(cmd.Flag("outputFormat").Value.String()),
		cmd.Flag("outputTo").Value.String(),
	)
	if errExport != nil {
		slog.Error(fmt.Sprintf("Failed to export document with error %s", errExport))
		os.Exit(1)
	}

	slog.Info("Finished processing input.")
}
