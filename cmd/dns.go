package cmd

import (
	"net"
	"os"
	"strconv"
	"strings"

	"github.com/DusanKasan/parsemail"
	"github.com/howellzach/eml-chckr/pkg/util"
	"github.com/spf13/cobra"
)

//TODO: Add spf checking for IP against sending domain
var (
	dnsCmd = &cobra.Command{
		Use:   "dns [eml file]",
		Short: "Gather DNS information about domains and IPs in an eml file",
		Long:  ``,
		Args: func(cmd *cobra.Command, args []string) error {
			if FileName == "" {
				if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
					return err
				}
			}
			return nil
		},
		Run: gatherDNS,
	}
)

func gatherDNS(ccmd *cobra.Command, args []string) {
	f, err := os.Open(FileName)
	util.CheckErr(err)
	email, err := parsemail.Parse(f)
	util.CheckErr(err)

	senderIP, senderDomain := util.ExtractSenderInfo(email)
	reverseIP, _ := net.LookupAddr(senderIP)
	txtRecords, _ := net.LookupTXT(senderDomain)
	data := map[string]string{
		"Sending Domain": senderDomain,
		"Sender IP":      senderIP,
		"PTR":            reverseIP[0],
	}
	hasSPF := false
	for _, txt := range txtRecords {
		if strings.HasPrefix(txt, "v=spf") {
			data["SPF Record"] = txt
			hasSPF = true
		}
	}
	data["SPF Record Present"] = strconv.FormatBool(hasSPF)
	util.GenerateTable("DNS Output", FileName, data)
}
