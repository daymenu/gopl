package main

import (
	"fmt"
	"log"
	"net/rpc"
)

// Fen 分
type Fen int64

// Yuan 元
type Yuan float64

func main() {
	c, err := rpc.DialHTTP("tcp", "127.0.0.1:7890")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	f := Fen(400)
	var y Yuan
	err = c.Call("Money.ToYuan", &f, &y)
	if err != nil {
		log.Fatal("money error:", err)
	}
	fmt.Println("价格为", y, "元")
}
