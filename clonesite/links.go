package main

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"golang.org/x/net/html"
)

const (
	a      = "a"
	img    = "img"
	audio  = "audio"
	embed  = "embed"
	video  = "video"
	script = "script"
	link   = "link"
)

var linkTags = map[string]string{
	a:      "href",
	img:    "src",
	audio:  "src",
	embed:  "src",
	video:  "src",
	script: "src",
	link:   "href",
}

func GetHtml(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func ParseHtml(b []byte) (*html.Node, error) {
	return html.Parse(bytes.NewReader(b))
}

func Links(n *html.Node) (links []string) {
	if n.Type != html.ElementNode {
		return
	}
	return
}
