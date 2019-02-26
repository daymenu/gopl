package main

import "fmt"

func main() {
	ch := make(chan int, 1)
	fmt.Println(ch)
	ch <- 4
	fmt.Println(<-ch)
	close(ch)
}
