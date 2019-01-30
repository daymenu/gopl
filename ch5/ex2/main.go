package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

var elements map[string]int = make(map[string]int)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlink1:%v\n", err)
		os.Exit(1)
	}
	finds(doc)
	for k, v := range elements {
		fmt.Printf("%s:%d\n", k, v)
	}
}

func finds(n *html.Node) {
	if n == nil {
		return
	}
	if n.Type == html.ElementNode {
		elements[n.Data]++
	}
	finds(n.FirstChild)
	finds(n.NextSibling)
}
