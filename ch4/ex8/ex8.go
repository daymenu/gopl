package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func main() {
	counts := make(map[string]int)
	in := bufio.NewReader(os.Stdin)
	for {
		r, _, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount:%v\n", err)
			os.Exit(1)
		}

		switch {
		case unicode.IsLetter(r):
			counts["letter"]++
			fallthrough
		case unicode.IsLower(r):
			counts["lower"]++
		}
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
}
