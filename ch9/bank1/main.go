package main

import (
	"fmt"
	"time"
)

var deposits = make(chan int, 1)
var balances = make(chan int, 1)
var withdrawed = make(chan bool, 1)
var withdraws = make(chan int, 1)

func Deposit(amount int) {
	deposits <- amount
}

func Balance() int {
	return <-balances
}

func WhitDraw(amount int) bool {
	withdraws <- amount
	return <-withdrawed
}
func teller() {
	var balance int
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case amount := <-withdraws:
			if balance < amount {
				withdrawed <- false
			} else {
				balance -= amount
				withdrawed <- true
			}
		}
	}
}

func init() {
	go teller()
}

func main() {
	Deposit(100)
	time.Sleep(2 * time.Second)
	fmt.Println(Balance())
}
