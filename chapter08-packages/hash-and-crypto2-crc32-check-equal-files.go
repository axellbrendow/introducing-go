package main

import (
	"fmt"
	"hash/crc32"
	"io"
	"os"
)

func getHash(filename string) (uint32, error) {
	// open the file
	file, err := os.Open(filename)
	if err != nil {
		return 0, err
	}
	// remember to always close opened files
	defer file.Close()
	// create a hasher
	hasher := crc32.NewIEEE()
	// copy the file into the hasher
	// - copy takes (dst, src) and returns (bytesWritten, error)
	_, err = io.Copy(hasher, file)
	// we don't care about how many bytes were written, but we do want to
	// handle the error
	if err != nil {
		return 0, err
	}

	return hasher.Sum32(), nil
}

func main() {
	hash1, err := getHash("test1.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	hash2, err := getHash("test2.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("hash1:", hash1)
	fmt.Println("hash2:", hash2)
	fmt.Println("hash1 == hash2:", hash1 == hash2)
}
