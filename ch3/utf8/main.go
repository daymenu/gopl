package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	str := "hello, 世界"
	fmt.Println("str: byte ", len(str))
	fmt.Println("str: rune ", utf8.RuneCountInString(str))

	for i := 0; i < len(str); {
		r, s := utf8.DecodeRuneInString(str[i:])
		i += s
		fmt.Printf("%c, %d\n", r, s)
	}
}
