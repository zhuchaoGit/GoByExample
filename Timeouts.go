package main

import (
	"time"
	"fmt"
)

func main() {
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "message"
	}()
	select {
	case msg := <-c1:
		fmt.Println("message", msg)
	case <-time.After(1 * time.Second):
		fmt.Println("out time")
	}
	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "message"
	}()
	select {
	case msg := <-c2:
		fmt.Println("message", msg)
	case <-time.After(3 * time.Second):
		fmt.Println("out time")
	}
}
