package utils

import (
	"log/slog"
	"os"

	"github.com/amolofos/tradesor/pkg/models/models_logLevel"
	"github.com/spf13/cobra"
)

func DefaultCmds(cmd *cobra.Command, args []string) {
	setLogger(cmd, args)
}

func setLogger(cmd *cobra.Command, args []string) {
	opts := &slog.HandlerOptions{
		Level: models_logLevel.SlogLevel(
			models_logLevel.LogLevel(cmd.Flag("logLevel").Value.String()),
		),
	}

	handler := slog.NewJSONHandler(os.Stdout, opts)
	logger := slog.New(handler)
	slog.SetDefault(logger)
}
