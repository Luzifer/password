package cli

import (
	"encoding/json"
	"errors"
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
		RunE:  actionCmdGet,
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

func actionCmdGet(cmd *cobra.Command, _ []string) error {
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
				return fmt.Errorf("checking for password pwnage: %w", err)
			}
		}

		if err != nil {
			if errors.Is(err, pwd.ErrLengthTooLow) {
				return fmt.Errorf("the password has to be more than 4 characters long to meet the security considerations")
			}

			return fmt.Errorf("getting password: %w", err)
		}

		if !flags.CLI.JSON {
			fmt.Println(password) //nolint:forbidigo
			continue
		}

		hashes, err := hasher.GetHashMap(password)
		if err != nil {
			return fmt.Errorf("generating password hashes: %w", err)
		}
		hashes["password"] = password
		if err = json.NewEncoder(os.Stdout).Encode(hashes); err != nil {
			return fmt.Errorf("encoding JSON: %w", err)
		}
	}

	return nil
}
