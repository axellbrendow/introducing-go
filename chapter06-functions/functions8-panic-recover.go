package main

import "fmt"

func main() {
	defer func() { // defer makes this function execute even after panic
		str := recover()
		fmt.Println(str)
	}()
	panic("PANIC")
}
