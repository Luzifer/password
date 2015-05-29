package main

import (
	"github.com/Luzifer/password/lib"
	"github.com/spf13/cobra"
)

const (
	version = "1.3.0"
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
