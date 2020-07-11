package math

import "fmt"

func init() {
	fmt.Println("mypackage/math package initialized")
}

// Average is an exported function because it starts with a capital letter
func Average(xs []float64) float64 {
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))
}
