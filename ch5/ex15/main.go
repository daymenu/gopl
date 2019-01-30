package main

import (
	"fmt"
)

func main() {
	fmt.Println(max(1, 2, 3, 4))
	fmt.Println(max(1))
	fmt.Println(max())
	fmt.Println(min(1, 2, 3, 4))
	fmt.Println(min(1))
	fmt.Println(min())
}

func max(ints ...int) (int, error) {
	if len(ints) == 0 {
		return 0, fmt.Errorf("请至少输入一个数字")
	}
	var m int = ints[0]
	for _, i := range ints {
		if m < i {
			m = i
		}
	}
	return m, nil
}

func min(ints ...int) (int, error) {
	if len(ints) == 0 {
		return 0, fmt.Errorf("请至少输入一个数字")
	}
	var m int = ints[0]
	for _, i := range ints {
		if m > i {
			m = i
		}
	}
	return m, nil
}
