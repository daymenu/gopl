package main

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

func main() {
	strconv.Atoi("12")
	var i int8 = 1
	fmt.Println(i<<7 - 1)
	fmt.Println(len("hello, 世界"))
	fmt.Println(utf8.RuneCountInString("hello, 世界"))
	fmt.Println(string(1234567))
	fmt.Printf("%s:%[1]T\n", fmt.Sprintf("%t", true))
	fmt.Printf("%s:%[1]T\n", fmt.Sprintf("%d", 17))
	fmt.Printf("%s:%[1]T\n", fmt.Sprintf("%.2f", 17.777))
	fmt.Printf("%s:%[1]T\n", strconv.FormatBool(true))
	fmt.Printf("%s:%[1]T\n", strconv.FormatInt(17, 8))
	fmt.Printf("%s:%[1]T\n", strconv.FormatFloat(17, byte('f'), 0, 64))
}

func Btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

func Itob(i int) bool {
	return i == 1
}
