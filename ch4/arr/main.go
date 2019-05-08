package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println(utf8.UTFMax)
	var arr [3]int
	fmt.Println(arr[0])
	fmt.Println(arr[len(arr)-1])

	for i, v := range arr {
		fmt.Printf("%d %d\n", i, v)
	}

	for _, v := range arr {
		fmt.Printf("%d\n", v)
	}

	var arr1 = [...]int{1, 2, 3}
	fmt.Println(arr1)

	var arr2 = [...]int{2: 22, 3: 33, 4: 44}
	fmt.Println(arr2)
}
