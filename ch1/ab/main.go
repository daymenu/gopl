package main

import "fmt"

func main() {
	a, b := 3, 2
	fmt.Printf("%032b %032b", a, b)
	fmt.Printf("%032b %032b\n", a^b, (a^b)^b)
	fmt.Printf("%032b %032b\n", a^b, (a^b)^a)
}
