package cmd

import (
	"fmt"
	"os"

	"amolofos/tradesor/cmd/transform"

	"github.com/spf13/cobra"
)

var (
	verbose bool   = false
	catalog string = ""
)

var rootCmd = &cobra.Command{
	Use:   "tradesor",
	Short: "tradesor - a simple CLI to transform and tradesor xml data",
	Long:  "tradesor is a simple CLI to transform and tradesor xml data",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
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
