package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	words := make(map[string]int)
	wordcount := 0
	input := bufio.NewScanner(os.Stdin)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		line := input.Text()
		words[line]++
		wordcount++
	}

	for w, n := range words {
		fmt.Printf("%s:%f\n", w, float64(n)/float64(wordcount))
	}
}
