package main

import (
	"fmt"
	"os"
)

func main() {
	for index, value := range os.Args {
		fmt.Println(index, " ", value)
	}
}
