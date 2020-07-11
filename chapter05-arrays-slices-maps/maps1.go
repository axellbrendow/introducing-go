package main

import "fmt"

func main() {
	// var x map[string]int // Here 'x' would be nil
	x := make(map[string]int)
	x["key"] = 10
	fmt.Println(x["key"])
	delete(x, "key")
	fmt.Println(x)
}
