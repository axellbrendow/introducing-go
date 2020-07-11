package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "from 1"
			time.Sleep(time.Second * 2)
		}
	}()

	go func() {
		for {
			c2 <- "from 2"
			time.Sleep(time.Second * 3)
		}
	}()

	go func() {
		for {
			/*
				select picks the first channel that is ready and receives from it
				(or sends to it).
			*/
			select {
			case msg1 := <-c1:
				fmt.Println(msg1)
			case msg2 := <-c2:
				fmt.Println(msg2)
			case <-time.After(time.Second):
				/*
					time.After creates a channel, and after the given duration, will
					send the current time on it (we weren’t interested in the time,
					so we didn’t store it in a variable).
				*/
				fmt.Println("timeout")
			default:
				fmt.Println("nothing ready")
				time.Sleep(time.Second)
			}
		}
	}()

	var input string
	fmt.Scanln(&input)
}
