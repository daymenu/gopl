package main

import "fmt"

func main() {
	defer func() {
		p := recover()
		fmt.Println(p)
	}()
	panicint(5)
}

func panicint(n int) {
	panic(n)
}
