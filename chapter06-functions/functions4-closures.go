package main

import "fmt"

func main() {
	x := 0

	addToX := func(args ...int) {
		for _, v := range args {
			x += v
		}
	}

	addToX(1)

	fmt.Println(x)
}
