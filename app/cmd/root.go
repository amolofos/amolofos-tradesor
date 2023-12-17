package cmd

import (
	"fmt"
	"os"

	"github.com/amolofos/tradesor/cmd/transform"
	"github.com/amolofos/tradesor/cmd/utils"
	"github.com/amolofos/tradesor/pkg/models/models_logLevel"

	"github.com/spf13/cobra"
)

var (
	logLevel models_logLevel.LogLevel = models_logLevel.Info
	catalog  string                   = ""
)

var rootCmd = &cobra.Command{
	Use:   "tradesor",
	Short: "tradesor - a simple CLI to transform and tradesor xml data",
	Long:  "tradesor is a simple CLI to transform and tradesor xml data",
	Run:   root,
}

func init() {
	rootCmd.PersistentFlags().VarP(&logLevel, "logLevel", "", "What log level to user to use: "+models_logLevel.GetAllSupportedValues())
	rootCmd.PersistentFlags().StringVarP(&catalog, "catalog", "c", "", "Location of the catalog. It can be a url or a local file.")
	rootCmd.MarkFlagRequired("catalog")

	rootCmd.AddCommand(transform.TransformCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Whoops. There was an error while executing your CLI '%s'", err)
		os.Exit(1)
	}
}

func root(cmd *cobra.Command, args []string) {
	utils.DefaultCmds(cmd, args)

	if len(args) == 0 {
		cmd.Help()
		os.Exit(0)
	}
}
