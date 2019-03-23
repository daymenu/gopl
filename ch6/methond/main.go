package main

import (
	"fmt"
	"math"
)

type Point struct{ x, y int }

func main() {
	p := Point{1, 4}
	q := Point{1, 1}
	fmt.Println(p.Distance(q))
}

func (p Point) Distance(q Point) float64 {
	x := math.Pow(float64(p.x-q.x), 2.0) + math.Pow(float64(p.y-q.y), 2.0)
	return math.Sqrt(x)
}
