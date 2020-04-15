package main

import "fmt"

func main() {
	x := [5]float64{ // number of elements is required
		98,
		3,
		7,
		2,
		83, // trailing comma is required
	}

	var total float64 = 0

	for _, value := range x { // Here, '_' would be the index (iterator)
		total += value
	}

	fmt.Println(total / float64(len(x)))
}
