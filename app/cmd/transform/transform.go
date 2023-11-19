package transform

import (
	"log/slog"

	"github.com/spf13/cobra"

	modelFormat "amolofos/tradesor/pkg/models/outputFormat"
	"amolofos/tradesor/pkg/services/loader"
	"amolofos/tradesor/pkg/services/transformer"
	"amolofos/tradesor/pkg/services/unloader"
)

var (
	outputTo     string = ""
	outputFormat modelFormat.OutputFormat
)

var TransformCmd = &cobra.Command{
	Use:   "transform [--outputFormat facebook]",
	Short: "transform Tradesor catalog",
	Long:  "transform Tradesor catalog",
	Run:   transform,
}

func init() {
	TransformCmd.Flags().StringVarP(&outputTo, "outputTo", "", "./output", "Location of the output files. It can be a url (for wordpress) or a local directory.")
	TransformCmd.Flags().VarP(&outputFormat, "outputFormat", "", "What format to use for the output: "+modelFormat.GetAllSupportedValues())
}

func transform(cmd *cobra.Command, args []string) {

	slog.Info(cmd.Flag("catalog").Value.String())
	slog.Info(cmd.Flag("outputFormat").Value.String())
	slog.Info(cmd.Flag("outputTo").Value.String())

	var l = &loader.Loader{}
	l.Init()

	var t = &transformer.Transformer{}
	t.Init()

	var u = &unloader.Unloader{}
	u.Init()

	doc, errLoad := l.Load(cmd.Flag("catalog").Value.String())
	if errLoad != nil {
		slog.Error("Failed to load document with error ", errLoad)
	}

	out, errTransform := t.Transform(doc, modelFormat.OutputFormat(cmd.Flag("outputFormat").Value.String()))
	if errTransform != nil {
		slog.Error("Failed to transform document with error ", errTransform)
	}

	errUnload := u.Unload(out, cmd.Flag("outputTo").Value.String())
	if errUnload != nil {
		slog.Error("Failed to unload document with error ", errUnload)
	}
}
