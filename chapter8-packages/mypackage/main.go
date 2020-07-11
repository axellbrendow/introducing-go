package main

import (
	"fmt"
	"mypackage/math"
)

// Enter in this folder and execute `go run main.go`

func init() {
	fmt.Println("mypackage package initialized")
}

func main() {
	fmt.Println()
	fmt.Println("Run `go doc mypackage/math` to see documentation")
	fmt.Println()
	xs := []float64{1, 2, 3, 4}
	avg := math.Average(xs)
	fmt.Println(avg)
}
