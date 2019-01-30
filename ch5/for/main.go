package main

import "fmt"

// 捕获迭代变量
func main() {
	var printSlice []func()
	for i := 1; i <= 10; i++ {
		i := i
		printSlice = append(
			printSlice, func() {
				fmt.Println(i)
				fmt.Println(&i)
			})
	}
	for _, f := range printSlice {
		f()
	}
}
