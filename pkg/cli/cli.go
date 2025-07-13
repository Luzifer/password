package cli

import "github.com/spf13/cobra"

var version = "dev"

// Execute is the main entrypoint for the CLI execution
func Execute() error {
	rootCmd := cobra.Command{
		Use:   "password",
		Short: "generates secure random passwords",
	}

	rootCmd.AddCommand(
		getCmdGet(),
		getCmdServe(),
		getCmdVersion(),
	)

	return rootCmd.Execute() //nolint:wrapcheck
}
