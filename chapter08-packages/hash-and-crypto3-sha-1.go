package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {
	hasher := sha1.New()
	hasher.Write([]byte("test"))
	hashBytes := hasher.Sum([]byte{})
	fmt.Println(hashBytes)
}
