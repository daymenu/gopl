package main

import (
	"fmt"
	"math"
)

type Point struct{ X, Y float64 }
type Path []Point

func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p Path) Distance() float64 {
	sum := 0.0
	for i := range p {
		if i < 1 {
			continue
		}
		sum += p[i-1].Distance(p[i])
	}
	return sum
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func main() {
	p := Point{1, 2}
	q := Point{4, 6}
	fmt.Println(Distance(p, q))
	fmt.Println(p.Distance(q))
	path := Path{p, q}
	fmt.Println(path.Distance())
	perim := Path{{1, 1}, {5, 1}, {5, 4}, {1, 1}}
	fmt.Println(perim.Distance()) // "1

	r := &Point{1, 2}
	r.ScaleBy(2)
	fmt.Println(*r)

	p.ScaleBy(3)
	fmt.Println(p)

	pptr := &Point{4, 4}
	pptr.ScaleBy(5)
	fmt.Println((*pptr).Distance(Point{3, 4}))
	fmt.Println(*pptr)
}
