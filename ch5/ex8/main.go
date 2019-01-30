package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func ElementByID(doc *html.Node, id string) *html.Node {
	return forEachNode(doc, id, startElement, endElement)
}

func main() {
	for _, url := range os.Args[1:] {
		doc, err := getDoc(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "findlink2: %v\n", err)
			continue
		}
		fmt.Println(ElementByID(doc, "result_logo"))
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

func startElement(n *html.Node, id string) bool {
	if n.Type == html.ElementNode {
		for _, v := range n.Attr {
			if v.Key == "id" && v.Val == id {
				return false
			}
		}
	}
	return true
}
func endElement(n *html.Node, id string) bool {
	return true
}
func forEachNode(n *html.Node, id string, pre, post func(n *html.Node, id string) bool) *html.Node {
	if pre != nil {
		find := pre(n, id)
		if !find {
			return n
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		find := forEachNode(c, id, pre, post)
		if find != nil {
			return find
		}
	}
	return nil
}
