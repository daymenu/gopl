package main

import (
	"fmt"
	"time"
)

func main() {
	go spinner(1 * time.Second)
	n := 4
	fmt.Printf("\rfib(%d)=%d", n, fib(n))
}
func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}
func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
