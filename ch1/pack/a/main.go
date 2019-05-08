package main

import (
	"fmt"

	"github.com/gopl/ch1/pack/b"
	"github.com/gopl/ch1/pack/c"
)

func main() {
	fmt.Println(b.Name)
	fmt.Println(c.Name)
}

func init() {
	fmt.Println("init main")
}
