package main

import "fmt"

func f() (int, int) {
	return 5, 6
}

func f2() (r int) { // Create a variable called r to be the return value
	r = 1
	return
}

func main() {
	x, y := f()
	fmt.Println(x, y)
	fmt.Println(f2())
}
