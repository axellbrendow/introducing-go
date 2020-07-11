package main

import "fmt"

func average(values []float64) float64 {
	total := 0.0
	for _, v := range values {
		total += v
	}
	return total / float64(len(values))
}

func main() {
	average := average([]float64{98, 93, 77, 82, 83})
	fmt.Println(average)
}
