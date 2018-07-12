package main

import "fmt"

func main() {
	messages := make(chan string, 1)
	signals := make(chan bool, 1)

	select {
	case msg := <-messages:
		fmt.Println("receieved message", msg)
	default:
		fmt.Println("no message")
	}

	msg := "hello"
	select {
	case messages <- msg: // buffered or no buffer
		fmt.Println("send hello")
	default:
		fmt.Println("no send")
	}
	select {
	case msg := <-messages:
		fmt.Println("recieved message", msg)
	case sig := <-signals:
		fmt.Println("recieved signal", sig)
	default:
		fmt.Println("nothing happened")
	}
}
