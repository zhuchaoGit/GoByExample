package main

import "fmt"

func main() {
	msgs := make(chan string, 2)
	go func() {
		var first, second string
		fmt.Scanln(&first, &second)
		msgs <- first
		msgs <- second
	}()
	fmt.Println(<-msgs)
	fmt.Println(<-msgs)
}
