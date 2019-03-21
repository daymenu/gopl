package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	counts := make(map[string]int)
	if len(os.Args[1:]) == 0 {
		countlines(os.Stdin, counts)
	} else {
		for _, fileName := range os.Args[1:] {
			f, err := os.Open(fileName)
			if err != nil {
				log.Fatal(err)
				continue
			}
			countlines(f, counts)
			f.Close()
		}
	}

	for line, sum := range counts {
		if sum > 1 {
			fmt.Println(line, ":", sum)
		}
	}
}

func countlines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	if input.Scan() {
		counts[input.Text()]++
	}
}
