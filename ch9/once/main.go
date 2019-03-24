package main

import (
	"fmt"
	"sync"
	"time"
)

var mu sync.Once
var stus map[string]int

func main() {
	fmt.Println(getStus("j"))
	go func() { fmt.Println(getStus("h")) }()
	time.Sleep(time.Second)
}

func setStus() {
	stus = map[string]int{"h": 1, "j": 2}
}

func getStus(name string) int {
	if stus == nil {
		mu.Do(setStus)
	}
	return stus[name]
}
