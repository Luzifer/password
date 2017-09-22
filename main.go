package main

import (
	"github.com/Luzifer/password/lib"
	"github.com/spf13/cobra"
)

var (
	version = "dev"
)

var pwd *securepassword.SecurePassword

func init() {
	pwd = securepassword.NewSecurePassword()
}

func main() {
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
