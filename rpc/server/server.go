package main

import (
	"fmt"
	"net/http"
	"net/rpc"
)

// Money 钱
type Money int64

// Fen 分
type Fen int64

// Yuan 元
type Yuan float64

func rpcMoney() {
	money := new(Money)
	fmt.Println("money===", money)
	rpc.Register(money)
	rpc.HandleHTTP()
	err := http.ListenAndServe(":7890", nil)
	if err != nil {
		fmt.Println("rpc error:", err)
	}
}

// ToYuan 转化成元
func (m *Money) ToYuan(f *Fen, y *Yuan) error {
	*y = Yuan(float64(*f) / 100)
	return nil
}

// ToFen 转化成分
func (m *Money) ToFen(f *Yuan, y *Fen) error {
	*y = Fen(float64(*f) * 100)
	return nil
}

func main() {
	rpcMoney()
}
