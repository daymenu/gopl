package main

import (
	"fmt"
	"time"

	"github.com/gopl/ch9/bank"
)

func main() {
	go func() {
		bank.Deposit(200)           //A1
		fmt.Println(bank.Balance()) //A2
	}()
	go bank.Deposit(100) //B
	time.Sleep(1 * time.Second)
	fmt.Println(bank.Balance())
}
