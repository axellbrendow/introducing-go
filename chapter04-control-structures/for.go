package main

import "fmt"

func for1() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func for2() {
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i = i + 1
	}
}

func main() {
	fmt.Println("First for style")
	for1()
	fmt.Println("\nSecond for style")
	for2()
}
