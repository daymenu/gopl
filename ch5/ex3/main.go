package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

var noSeen map[string]bool = map[string]bool{"script": true, "style": true}

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
		<script>alert(1)</script>
	</body>
	</html>`))
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlink1:%v\n", err)
		os.Exit(1)
	}
	showText(doc)
}

func showText(n *html.Node) {
	if n == nil {
		return
	}

	if n.Type == html.TextNode && !noSeen[n.Parent.Data] {
		fmt.Println(n.Data)
	}
	showText(n.FirstChild)
	showText(n.NextSibling)
}
