package main

import "fmt"

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5))
	fmt.Println(sum())
	fmt.Println(sum(1))

	mints := []int{1, 2, 3, 4, 5}
	fmt.Println(sum(mints...))
}

func sum(vals ...int) int {
	s := 0
	for _, val := range vals {
		s += val
	}
	return s
}
