package main

import "fmt"

func main() {
	msg := make(chan string)
	go func() {
		var s string
		if n, err := fmt.Scanln(&s); err != nil {
			fmt.Println(n, err)
		}
		msg <- s
	}()

	str := <-msg
	fmt.Println(str)
}
