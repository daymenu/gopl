package main

import (
	"fmt"
	"log"
)

func main() {
	f(3)
}

func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x)
	defer func() {
		if p := recover(); p != nil {
			log.Fatal(p)
		}
	}()
	f(x - 1)
}
