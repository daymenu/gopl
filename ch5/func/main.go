package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func main() {
	var f func(int, int) int

	f = add

	fmt.Println(f(1, 2))
	fmt.Println(sum(1, 2, 3))
}

func sum(nums ...int) int {
	var s int
	for _, num := range nums {
		s += num
	}
	return s
}
