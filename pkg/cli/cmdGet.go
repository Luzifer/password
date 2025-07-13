package cli

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	pwd "github.com/Luzifer/password/lib/v2"
	"github.com/Luzifer/password/v2/pkg/hasher"
)

const (
	defaultPasswordCount  = 1
	defaultPasswordLength = 20
)

func getCmdGet() *cobra.Command {
	cmd := cobra.Command{
		Use:   "get",
		Short: "generate and return a secure random password",
		Run:   actionCmdGet,
	}

	cmd.Flags().BoolVarP(&flags.CLI.JSON, "json", "j", false, "return output in JSON format")
	cmd.Flags().IntVarP(&flags.CLI.Length, "length", "l", defaultPasswordLength, "length of the generated password")
	cmd.Flags().IntVarP(&flags.CLI.Num, "number", "n", defaultPasswordCount, "number of passwords to generate")
	cmd.Flags().BoolVarP(&flags.CLI.SpecialCharacters, "special", "s", false, "use special characters in your password")

	cmd.Flags().BoolVarP(&flags.CLI.XKCD, "xkcd", "x", false, "use XKCD style password")
	cmd.Flags().StringVar(&flags.CLI.Separator, "separator", "", "add separator between words of XKCD style password")
	cmd.Flags().BoolVarP(&flags.CLI.PrependDate, "date", "d", true, "prepend current date to XKCD style passwords")

	cmd.Flags().Bool("check-hibp", false, "Check HaveIBeenPwned for this password")

	return &cmd
}

func actionCmdGet(cmd *cobra.Command, args []string) {
	var (
		password string
		err      error
	)

	pwd.DefaultXKCD.Separator = flags.CLI.Separator

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
				fmt.Println("An unknown error occurred")
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
