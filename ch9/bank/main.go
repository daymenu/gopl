package main

import "fmt"

var balance int = 0

func main() {
	var x []int
	go func() { x = make([]int, 10) }()
	go func() { x = make([]int, 10000) }()
	x[9999] = 1
}

func deposit(deposit int) {
	balance += deposit
}

func getBalance() int {
	return balance
}

func bank() {
	ch := make(chan int, 1)
	go func() {
		deposit(200)
		ch <- getBalance()
	}()

	go deposit(100)
	fmt.Println(<-ch)
}
