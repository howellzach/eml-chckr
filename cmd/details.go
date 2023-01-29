package cmd

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"net/mail"
	"os"
	"strconv"
	"strings"

	"github.com/DusanKasan/parsemail"
	"github.com/howellzach/eml-chckr/pkg/util"
	"github.com/spf13/cobra"
)

type attachment struct {
	filename    string
	contenttype string
	sha256hash  string
	sha1hash    string
	md5hash     string
}

var (
	detailsCmd = &cobra.Command{
		Use:   "details [eml file]",
		Short: "Print general information about an eml file",
		Long:  ``,
		Args: func(cmd *cobra.Command, args []string) error {
			// Optionally run one of the validators provided by cobra
			if FileName == "" {
				if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
					return err
				}
			}
			return nil
		},
		Run: details,
	}
)

func details(ccmd *cobra.Command, args []string) {
	f, err := os.Open(FileName)
	util.CheckErr(err)
	email, err := parsemail.Parse(f)
	util.CheckErr(err)
	senderIP, senderDomain := util.ExtractSenderInfo(email)
	fromList := createEmlAddrList(email.From)
	replyToList := createEmlAddrList(email.ReplyTo)
	toList := createEmlAddrList(email.To)
	ccList := createEmlAddrList(email.Cc)

	data := map[string]string{
		"Date":          email.Date.Local().String(),
		"From":          strings.Join(fromList, ", "),
		"Reply to":      strings.Join(replyToList, ", "),
		"Return-Path":   email.Header.Get("return-path"),
		"To":            strings.Join(toList, ", "),
		"CCs":           strings.Join(ccList, ", "),
		"Subject":       email.Subject,
		"Sender IP":     senderIP,
		"Sender Domain": senderDomain,
	}

	var attachments []attachment
	if len(email.Attachments) > 0 {
		for i, a := range email.Attachments {
			data[fmt.Sprintf("Attachment %d", i+1)] = a.Filename
			attachments = append(attachments, createAttachment(a))
		}
	} else {
		data["Email has no attachments"] = "N/A"
	}
	util.GenerateTable("Details Output", FileName, data)

	if len(attachments) > 0 {
		outputAttachmentTables(attachments)
	}
}

func createEmlAddrList(m []*mail.Address) []string {
	var addrList []string
	for _, addr := range m {
		addrList = append(addrList, fmt.Sprintf("%s <%s>", addr.Name, addr.Address))
	}
	return addrList
}

func createAttachment(a parsemail.Attachment) attachment {
	h256 := sha256.New()
	_, err := io.Copy(h256, a.Data)
	util.CheckErr(err)
	hsha1 := sha1.New()
	_, err = io.Copy(hsha1, a.Data)
	util.CheckErr(err)
	hmd5 := md5.New()
	_, err = io.Copy(hmd5, a.Data)
	util.CheckErr(err)
	return attachment{
		filename:    a.Filename,
		contenttype: a.ContentType,
		sha256hash:  hex.EncodeToString(h256.Sum(nil)),
		sha1hash:    hex.EncodeToString(hsha1.Sum(nil)),
		md5hash:     hex.EncodeToString(hmd5.Sum(nil)),
	}
}

func outputAttachmentTables(attachments []attachment) {
	for i, a := range attachments {
		attachData := map[string]string{
			"Filename":  a.filename,
			"File type": a.contenttype,
			"SHA256":    a.sha256hash,
			"SHA1":      a.sha1hash,
			"MD5":       a.md5hash,
		}
		util.GenerateTable("Attachment Info", strconv.Itoa(i+1), attachData)
	}
}
