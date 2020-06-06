package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, r float64
}

// You can use dot syntax to access Circle and *Circle
func circleArea(c *Circle) float64 {
	return math.Pi * c.r * c.r
}

func main() {
	c := Circle{0, 0, 5}
	fmt.Println(circleArea(&c))
}
