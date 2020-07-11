package main

import "fmt"

func main() {
	var x []float64            // This is an slice
	x = make([]float64, 5)     // Slice of length 5 associated with an array of length 5
	x = make([]float64, 5, 10) // Slice of length 5 associated with an array of length 10

	arr := [5]float64{1, 2, 3, 4, 5}

	x = arr[1:4] // [2, 3, 4]
	fmt.Println(x)

	x = arr[:4] // [1, 2, 3, 4]
	fmt.Println(x)

	x = arr[1:] // [2, 3, 4, 5]
	fmt.Println(x)

	x = arr[:] // [1, 2, 3, 4, 5]
	fmt.Println(x)
}
