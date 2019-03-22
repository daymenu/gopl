package main

import "fmt"

type Point struct {
	X int
	Y int
}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle Circle
	Spokes int
}

func main() {

	w := Wheel{Circle{Point{8, 8}, 5}, 20}
	fmt.Println(w)

	w.Spokes = 50
	fmt.Println(w)

	w.Circle.Point.X = 9
	fmt.Println(w)
}
