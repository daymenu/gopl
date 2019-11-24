package main

import (
	"encoding/xml"
	"fmt"
	"strings"
)

func main() {
	type wei struct {
		XMLName     xml.Name `xml:"error" json:"-"`
		Ret         int      `xml:"ret" json:"-"`
		Message     string   `xml:"message" json:"-"`
		Skey        string   `xml:"skey" json:"Skey"`
		Wxsid       string   `xml:"wxsid" json:"Sid"`
		Wxuin       int64    `xml:"wxuin" json:"Uin"`
		PassTicket  string   `xml:"pass_ticket" json:"-"`
		DeviceID    string   `xml:"-" json:"DeviceID"`
		IsGrayScale int      `xml:"isgrayscale" json:"-"`
	}
	v := new(wei)
	data := `<error><ret>1203</ret><message></message></error>`
	err := xml.NewDecoder(strings.NewReader(data)).Decode(&v)
	fmt.Printf("%v %#v", err, v)
}
