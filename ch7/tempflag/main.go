package main

import (
	"flag"
	"fmt"

	"github.com/gopl/ch7/tempconv"
)

var temp = tempconv.CelsiusFlag("temp1", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
