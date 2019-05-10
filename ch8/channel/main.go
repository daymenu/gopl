package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var names = []string{"hanjian", "liliting", "taina"}
	echoName(names)
	time.Sleep(1 * time.Second)
}

func echoName(names []string) {
	var ch = make(chan string)
	var wait sync.WaitGroup
	for _, name := range names {
		wait.Add(1)
		go func(name string) {
			defer wait.Done()
			fmt.Println("go1 ", name)
			ch <- name
			fmt.Println("go1 ", name)
		}(name)
	}
	go func() {
		wait.Wait()
		close(ch)
	}()
	for n := range ch {
		fmt.Println("d:", n)
	}
	return
}
