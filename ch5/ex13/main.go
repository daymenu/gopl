package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gopl/ch5/links"
)

func main() {
	breadthFirst(crawl, os.Args[1:])
}

func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func crawl(link string) []string {
	fmt.Println(link)
	list, err := links.Extract(link)
	if err != nil {
		log.Print(err)
	}
	return list
}
