package main

import (
	"fmt"
	"time"
)

const noDely time.Duration = 0
const timeout = 5 * time.Minute

const (
	a = 1
	b
	e = iota
	f
	g
	c = 3
	d
)

func main() {
	fmt.Printf("%T %[1]v\n", noDely)
	fmt.Printf("%T %[1]v\n", timeout)
	fmt.Printf("%T %[1]v\n", time.Minute)
	fmt.Println(a, b, c, d, e, f, g)
}
