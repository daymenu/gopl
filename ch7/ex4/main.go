package main

import (
	"fmt"
	"io"

	"golang.org/x/net/html"
)

type StringReader string

func (s *StringReader) Read(p []byte) (n int, err error) {
	p = []byte(*s)
	return len(p), io.EOF
}
func NewReader(s string) io.Reader {
	sr := StringReader(s)
	return &sr
}

func Parse(s string) (*html.Node, error) {
	htmlReader := NewReader(s)
	return html.Parse(htmlReader)
}
func main() {
	htmldoc := `<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<meta http-equiv="X-UA-Compatible" content="ie=edge">
		<title>Document</title>
	</head>
	<body>
		
	</body>
	</html>`
	fmt.Println(Parse(htmldoc))
}
