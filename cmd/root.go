package cmd

import (
	"context"

	"github.com/elewis787/boa"
	"github.com/spf13/cobra"
)

func Execute() error {
	rootCmd := &cobra.Command{
		Version: "v0.0.1",
		Use:     "spoctify",
		Long:    "A Terminal UI for people who want to listen Spotify on boring black and white screen",
		Example: "spoclify",
	}

	rootCmd.SetHelpFunc(boa.HelpFunc)
	rootCmd.SetUsageFunc(boa.UsageFunc)

	rootCmd.AddCommand(AuthCommand())

	return rootCmd.ExecuteContext(context.Background())
}
