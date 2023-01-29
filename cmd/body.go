package cmd

import (
	"fmt"
	"os"

	"github.com/DusanKasan/parsemail"
	"github.com/howellzach/eml-chckr/pkg/util"
	"github.com/spf13/cobra"
)

var (
	bodyCmd = &cobra.Command{
		Use:   "body [eml file]",
		Short: "Print eml file message body",
		Long:  ``,
		Args: func(cmd *cobra.Command, args []string) error {
			if FileName == "" {
				if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
					return err
				}
			}
			return nil
		},
		Run: printBody,
	}
)

func printBody(ccmd *cobra.Command, args []string) {
	f, err := os.Open(FileName)
	util.CheckErr(err)
	email, err := parsemail.Parse(f)
	util.CheckErr(err)

	if HTMLBody {
		fmt.Println(email.HTMLBody)
	} else if len(email.TextBody) == 0 {
		fmt.Println("Email has no text body to print, try running command with the '--html' flag")
	} else {
		fmt.Println(email.TextBody)
	}
}
