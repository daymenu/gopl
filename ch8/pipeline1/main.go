package main

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go func() {
		for x := 0; ; x++ {
			naturals <- x
			if x > 8 {
				close(naturals)
				break
			}
		}
	}()
	go func() {
		for {
			x, ok := <-naturals
			if !ok {
				break
			}
			squares <- x * x
			close(squares)
		}
	}()

	for {
		if x, ok := <-squares; ok {
			fmt.Println(x)
		} else {
			break
		}
	}
}
