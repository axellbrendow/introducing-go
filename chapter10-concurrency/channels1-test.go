package main

import (
	"fmt"
	"time"
)

// `chan<-` restricts the direction of the channel `c`. In this case, `pinger()`
// can only send strings to the channel

func pinger(c chan<- string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

// `chan<-` restricts the direction of the channel `c`. In this case, `printer()`
// can only extract strings from the channel

func printer(c <-chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func ponger(c chan<- string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}

func main() {
	var c chan string = make(chan string)
	go pinger(c)
	go ponger(c)
	go printer(c)
	var input string
	fmt.Scanln(&input)
}
