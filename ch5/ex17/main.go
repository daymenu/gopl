package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func ElementByTagName(doc *html.Node, names ...string) []*html.Node {
	return forEachNode(doc, startElement, nil, names...)
}

func main() {
	for _, url := range os.Args[1:] {
		doc, err := getDoc(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "findlink2: %v\n", err)
			continue
		}
		for _, n := range ElementByTagName(doc, "div", "img") {
			fmt.Println(n.Data)
		}
	}
}

func getDoc(url string) (*html.Node, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s:%s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML:%v", url, err)
	}

	return doc, nil
}

func startElement(n *html.Node, names ...string) bool {
	if n.Type != html.ElementNode {
		return false
	}
	for _, name := range names {
		if n.Data == name {
			return true
		}
	}
	return false
}
func endElement(n *html.Node, id string) bool {
	return true
}
func forEachNode(n *html.Node, pre, post func(n *html.Node, name ...string) bool, names ...string) []*html.Node {
	var tagNodes []*html.Node
	if pre != nil {
		find := pre(n, names...)
		if find {
			tagNodes = append(tagNodes, n)
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		find := forEachNode(c, pre, nil, names...)
		if find != nil {
			tagNodes = append(tagNodes, find...)
		}
	}
	return tagNodes
}
