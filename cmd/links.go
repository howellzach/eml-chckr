package cmd

import (
	"os"
	"strings"

	"github.com/DusanKasan/parsemail"
	"github.com/howellzach/eml-chckr/pkg/util"
	"github.com/spf13/cobra"
)

var (
	linksCmd = &cobra.Command{
		Use:   "links [eml file]",
		Short: "Extract URLs from eml file",
		Long:  ``,
		Args: func(cmd *cobra.Command, args []string) error {
			if FileName == "" {
				if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
					return err
				}
			}
			return nil
		},
		Run: gatherLinks,
	}
)

func gatherLinks(ccmd *cobra.Command, args []string) {
	f, err := os.Open(FileName)
	util.CheckErr(err)
	email, err := parsemail.Parse(f)
	util.CheckErr(err)
	urls := util.ExtractURLs(email.HTMLBody)
	urls = removeBenignLinks(urls)
	data := make(map[string]string)
	var u string
	if len(urls) > 0 {
		for _, url := range urls {
			if UnescapeURLs {
				u = url
			} else {
				u = escapeURL(url)
			}
			data[u] = u
		}
	} else {
		data["Found no links in eml file"] = "Found no links in eml file"
	}
	util.GenerateTable("EML Links", FileName, data)
}

func removeBenignLinks(URLs []string) []string {
	var links []string
	benignLinks := []string{
		"http://schemas.microsoft.com/office/2004/12/omml",
		"http://www.w3.org/TR/REC-html40",
	}
	for _, url := range URLs {
		if stringNotInSlice(url, benignLinks) {
			links = append(links, url)
		}
	}
	return links
}

func stringNotInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return false
		}
	}
	return true
}

func escapeURL(u string) string {
	u = strings.Replace(u, "http", "hxxp", -1)
	u = strings.Replace(u, ":", "[:]", -1)
	u = strings.Replace(u, ".", "[.]", -1)
	return u
}
