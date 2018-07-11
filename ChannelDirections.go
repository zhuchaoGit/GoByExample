package main

import "fmt"

func ping(pings chan<- string, msg string) {
	pings <- msg
}
func pong(pings <-chan string, pongs chan<- string) {
	pongs <- <-pings
}

func main() {
	pings := make(chan string)
	pongs := make(chan string)
	ping(pings, "hello")
	pong(pings, pongs)
	fmt.Println(<-pongs)

}
