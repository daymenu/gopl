package main

//d
//d
import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := make(map[string][]string)
	if len(os.Args[1:]) == 0 {
		countlines(os.Stdin, counts, files)
	} else {
		for _, fileName := range os.Args[1:] {
			f, err := os.Open(fileName)
			if err != nil {
				log.Printf("du2:%v\n", err)
				continue
			}
			countlines(f, counts, files)
			f.Close()
		}
	}

	for line, sum := range counts {
		if sum > 1 {
			fmt.Println(files[line], "->", line, ":", sum)
		}
	}
}

func countlines(f *os.File, counts map[string]int, files map[string][]string) {
	fileName := f.Name()
	input := bufio.NewScanner(f)
	for input.Scan() {
		text := input.Text()
		counts[text]++
		files[text] = append(files[text], fileName)
	}
}
