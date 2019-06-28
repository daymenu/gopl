package main

import (
	"fmt"
	"gopl/ch5/links"
	"log"
	"os"
)

func main() {
	Crawl1()
	return
	worklists := make(chan []string)
	seenLinks := make(map[string]struct{})
	if len(os.Args[1:]) == 0 {
		log.Fatalln("Please input url")
	}
	go func() { worklists <- os.Args[1:] }()
	n := 1
	for ; n > 0; n-- {
		lists := <-worklists
		for _, link := range lists {
			if _, ok := seenLinks[link]; !ok {
				seenLinks[link] = struct{}{}
				n++
				go func(link string) {
					worklists <- crawl(link)
				}(link)
			}
		}
	}

	fmt.Println("crawl end")
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func Crawl1() {
	worklists := make(chan []string)
	seenLinks := make(map[string]struct{})
	unseenLinks := make(chan string)

	if len(os.Args[1:]) == 0 {
		log.Fatalln("Please input url")
	}
	go func() { worklists <- os.Args[1:] }()
	for n := 2; n > 0; n-- {
		go func() {
			for list := range unseenLinks {
				foundLinks := crawl(list)
				go func() {
					worklists <- foundLinks
				}()
			}
		}()
	}

	for lists := range worklists {
		for _, list := range lists {
			if _, ok := seenLinks[list]; !ok {
				fmt.Println("worklist:", list)
				seenLinks[list] = struct{}{}
				unseenLinks <- list
			}
		}
	}
}
