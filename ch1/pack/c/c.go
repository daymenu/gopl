package c

import "fmt"

var Name = "c"

func init() {
	Name = "c new"
	fmt.Println("init c")
}
