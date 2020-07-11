package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		fmt.Println("Could not open file")
		fmt.Println(err)
		return
	}
	defer file.Close()

	// get the file size
	stat, err := file.Stat()
	if err != nil {
		fmt.Println("Could not get file attributes")
		fmt.Println(err)
		return
	}

	// read the file
	fileBytes := make([]byte, stat.Size())
	_, err = file.Read(fileBytes)
	if err != nil {
		fmt.Println("Could not read file data")
		fmt.Println(err)
		return
	}

	str := string(fileBytes)
	fmt.Println(str)
}
