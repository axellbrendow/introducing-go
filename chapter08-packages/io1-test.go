package main

import (
	"bytes"
	"fmt"
)

func main() {
	var buf bytes.Buffer
	buf.Write([]byte("abc"))
	fmt.Println(buf.Bytes())
}
