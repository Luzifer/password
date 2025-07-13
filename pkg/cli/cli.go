package cli

import "github.com/spf13/cobra"

var version = "dev"

func Execute() {
	rootCmd := cobra.Command{
		Use:   "password",
		Short: "generates secure random passwords",
	}

	rootCmd.AddCommand(
		getCmdGet(),
		getCmdServe(),
		getCmdVersion(),
	)

	rootCmd.Execute()
}
