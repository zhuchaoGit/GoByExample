package main

import (
	"time"
	"fmt"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	go func() {
		time.Sleep(time.Second)
		c1 <- "c1"
	}()
	go func() {
		time.Sleep(time.Second)
		c2 <- "c2"
	}()
	for _ = range []int{1, 2} {
		select {
		case msg1 := <-c1:
			fmt.Println("c1", msg1)
		case msg2 := <-c2:
			fmt.Println("c1", msg2)
		}
	}
}
