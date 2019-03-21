package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep = "", ""
	for _, arg := range os.Args {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
