package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}
type Point struct{ x, y int }

type Path struct {
	start Point
	end   Point
}

type Circle struct {
	Point
	Radius int
}

func main() {
	var dilbert Employee
	dilbert.Address = "中国"
	fmt.Println(dilbert)
	var p Point
	p.x = 3
	p.y = 4
	fmt.Println(p)

	var c Circle
	c.Point.x = 1
	c.Point.y = 2
	c.Radius = 4
	fmt.Println(c)

	c1 := Circle{Point: Point{1, 2}, Radius: 3}
	fmt.Printf("%#v\n", c1)
}
