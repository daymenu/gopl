package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args[1:]) == 0 {
		fmt.Println("请输入你要转化的数字：")
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			text := input.Text()
			if text == "exit" {
				break
			}
			conv(text)
		}
	} else {
		for _, arg := range os.Args[1:] {
			conv(arg)
		}
	}
}

func conv(text string) {
	h, err := strconv.ParseFloat(text, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ex2:%v\n", err)
		os.Exit(1)
	}
	p := Pound(h)
	k := Kilo(h)
	fmt.Printf("%s = %s %s = %s\n", p, PToK(p), k, KToP(k))
}
