package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

var elements map[string]string = map[string]string{
	"a":      "href",
	"img":    "src",
	"script": "src",
	"link":   "href",
}

func main() {
	doc, err := html.Parse(strings.NewReader(`<html lang="en">
	<head>
		<meta charset="UTF-8"/>
		<meta name="viewport" content="width=device-width, initial-scale=1.0"/>
		<meta http-equiv="X-UA-Compatible" content="ie=edge"/>
		<title>Document</title>
	</head>
	<body>
		<a href = "666">666</a>
		<img src="777">
	</body>
	</html>`))
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlink1:%v\n", err)
		os.Exit(1)
	}
	fmt.Println(visit(nil, doc))
}

func visit(links []string, n *html.Node) []string {
	if n == nil {
		return links
	}
	attrStr, ok := elements[n.Data]
	if n.Type == html.ElementNode && ok {
		for _, a := range n.Attr {
			if a.Key == attrStr {
				links = append(links, a.Val)
			}
		}
	}
	links = visit(links, n.FirstChild)
	links = visit(links, n.NextSibling)
	return links
}
