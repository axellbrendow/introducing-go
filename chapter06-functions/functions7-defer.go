package main

import (
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("functions7-defer.go")
	// What a beautiful use of defer
	defer f.Close()                           // This will be the last call
	defer fmt.Println("This is the end1....") // This is the penultimate
	defer fmt.Println("This is the end2....") // This is the antepenultimate (the first call)
	fmt.Println("f.Close() will be moved to the end of main()")
}
