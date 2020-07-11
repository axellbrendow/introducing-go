package main

import (
	"fmt"
	"os"
)

func main() {
	dir, err := os.Open(".")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer dir.Close()

	// -1 would be the limit of entries to get
	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, fi := range fileInfos {
		fmt.Println(fi.Name())
	}
}
