package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

//TODO: Add command for extracting links, domains, IP addresses from email body

var version = "0.0.1"
var FileName string
var HTMLBody bool
var UnescapeURLs bool

var (
	rootCmd = &cobra.Command{
		Use:     "eml-chckr",
		Short:   "eml-chckr – a simple eml file analysis tool",
		Version: version,
		Long: `eml-chckr – eml file analysis tool
		
  An 'EML_FILE' environment variable can be set rather than specifying a specific eml file as an input to the commands`,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if FileName == "" || len(args) > 0 {
				FileName = args[0]
			}
		},
	}
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)
	bodyCmd.PersistentFlags().BoolVar(&HTMLBody, "html", false, "Print the full HTML body for eml file")
	linksCmd.PersistentFlags().BoolVarP(&UnescapeURLs, "unescape", "u", false, "Unescape URLs that are gathered from eml file")
	rootCmd.AddCommand(detailsCmd)
	rootCmd.AddCommand(dnsCmd)
	rootCmd.AddCommand(bodyCmd)
	rootCmd.AddCommand(linksCmd)
}

func initConfig() {
	viper.AutomaticEnv()
	FileName = viper.GetString("EML_FILE")
}
