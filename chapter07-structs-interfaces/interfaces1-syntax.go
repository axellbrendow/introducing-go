package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, r float64
}

// You can use dot syntax to access Circle and *Circle
func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

type Shape interface {
	area() float64
}

// Using an interface we can make generic functions like this
func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

func main() {
	fmt.Println(totalArea(&Circle{x: 0, y: 0, r: 5},
		&Rectangle{x1: 0, y1: 0, x2: 10, y2: 10}))
}
