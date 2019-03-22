package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "我爱你"
	fmt.Println(len(s))
	fmt.Println(s[3:])

	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d %c\n", i, r)
		i += size
	}

	for i, r := range s {
		fmt.Printf("%d %c\n", i, r)
	}
}
