package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/Luzifer/password/hasher"
	pwd "github.com/Luzifer/password/lib"
)

func getCmdGet() *cobra.Command {
	cmd := cobra.Command{
		Use:   "get",
		Short: "generate and return a secure random password",
		Run:   actionCmdGet,
	}

	cmd.Flags().BoolVarP(&flags.CLI.JSON, "json", "j", false, "return output in JSON format")
	cmd.Flags().IntVarP(&flags.CLI.Length, "length", "l", 20, "length of the generated password")
	cmd.Flags().IntVarP(&flags.CLI.Num, "number", "n", 1, "number of passwords to generate")
	cmd.Flags().BoolVarP(&flags.CLI.SpecialCharacters, "special", "s", false, "use special characters in your password")

	cmd.Flags().BoolVarP(&flags.CLI.XKCD, "xkcd", "x", false, "use XKCD style password")
	cmd.Flags().BoolVarP(&flags.CLI.PrependDate, "date", "d", true, "prepend current date to XKCD style passwords")

	cmd.Flags().Bool("check-hibp", false, "Check HaveIBeenPwned for this password")

	return &cmd
}

func actionCmdGet(cmd *cobra.Command, args []string) {
	var (
		password string
		err      error
	)

	for i := 0; i < flags.CLI.Num; i++ {

	regenerate:
		if flags.CLI.XKCD {
			password, err = pwd.DefaultXKCD.GeneratePassword(flags.CLI.Length, flags.CLI.PrependDate)
		} else {
			password, err = pwd.NewSecurePassword().GeneratePassword(flags.CLI.Length, flags.CLI.SpecialCharacters)
		}

		if c, _ := cmd.Flags().GetBool("check-hibp"); c {
			switch pwd.CheckHIBPPasswordHash(password) {
			case pwd.ErrPasswordInBreach:
				goto regenerate
			case nil:
				// Just do nothing
			default:
				fmt.Printf("Unable to check for password pwnage: %s", err)
				os.Exit(1)
			}
		}

		if err != nil {
			switch {
			case err == pwd.ErrLengthTooLow:
				fmt.Println("The password has to be more than 4 characters long to meet the security considerations")
			default:
				fmt.Println("An unknown error occured")
			}
			os.Exit(1)
		}

		if !flags.CLI.JSON {
			fmt.Println(password)
			continue
		}

		hashes, err := hasher.GetHashMap(password)
		if err != nil {
			fmt.Printf("Unable to generate hashes: %s", err)
			os.Exit(1)
		}
		hashes["password"] = password
		json.NewEncoder(os.Stdout).Encode(hashes)

	}
}
