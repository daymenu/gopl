package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	var w io.Writer
	w = os.Stdout
	f := w.(*os.File)
	c, ok := w.(*bytes.Buffer)
	fmt.Printf("%T\n", f)
	fmt.Printf("%T %t\n", c, ok)

	_, err := os.Open("/no/such/file")
	fmt.Println(err)
	fmt.Printf("%#v\n", err)
	var xx interface{}
	xx = 3.8
	switch x := xx.(type) {
	case int:
		fmt.Println("---int---")
	default:
		fmt.Printf("%T", x)
	}
}
