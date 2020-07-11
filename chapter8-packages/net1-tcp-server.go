package main

import (
	"encoding/gob"
	"fmt"
	"net"
	"time"
)

func handleServerConnection(connection net.Conn) {
	// receive the message
	var msg string
	err := gob.NewDecoder(connection).Decode(&msg)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Received '%s'\n", msg)
	}
	connection.Write([]byte("Data received!"))
	connection.Close()
}

func server() {
	// listen on a port
	listener, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		// accept a connection
		connection, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		// handle the connection
		go handleServerConnection(connection)
	}
}

func client() {
	// connect to the server
	connection, err := net.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer connection.Close()

	msg := "Hello, World"
	fmt.Printf("Sending '%s'\n", msg)

	// send the message
	err = gob.NewEncoder(connection).Encode(msg)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	go server()
	time.Sleep(200 * time.Millisecond) // Just to wait server start
	go client()
	var input string
	fmt.Scanln(&input)
}
