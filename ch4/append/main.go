package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(nums)
	fmt.Println(appendInt(nums, 6))
}

func appendInt(x []int, y int) []int {
	var z []int = x
	zlen := len(x) + 1
	if zlen > cap(x) {
		z = make([]int, zlen, 2*zlen)
	} else {
		z = x[:zlen]
	}
	copy(z, x)
	z[len(x)] = y
	return z
}
