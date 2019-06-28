package main

import (
	"log"
	"os"

	"gopl/ch5/links"
)

func main() {
	worklist := make(chan []string)
	unseenLinks := make(chan string)

	go func() { worklist <- os.Args[1:] }()
	for i := 0; i < 2; i++ {
		go func(i int) {
			for link := range unseenLinks {
				foundLinks := crawl(link)
				go func() {
					if len(foundLinks) == 0 {
						close(worklist)
					} else {
						worklist <- foundLinks
					}
				}()
			}
		}(i)
	}

	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link
			}
		}
	}
}

func crawl(url string) []string {
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}
