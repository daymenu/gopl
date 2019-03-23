package main

import (
	"encoding/json"
	"fmt"
	"io"
)

type Point struct{ X, Y int }

func main() {
	p := Point{1, 2}
	pJson, _ := json.MarshalIndent(p, " ", "\t")
	var p2 Point
	json.Unmarshal(pJson, &p2)
	fmt.Println(p2)
	fmt.Printf("%#v\n", io.EOF)

}
