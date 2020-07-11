package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	// read the file
	fileBytes, err := ioutil.ReadFile("test.txt")
	if err != nil {
		fmt.Println("Could not read file data")
		fmt.Println(err)
		return
	}

	str := string(fileBytes)
	fmt.Println(str)
}
