package main

import "fmt"

func f(str string) {
	for i := 0; i < 3; i++ {
		fmt.Println(str, ":", i)
	}
}
func main() {
	f("main")

	go f("routine")
	go f("routine2")
	go func(str string) {
		fmt.Println(str)
	}("anonymous")

	fmt.Scanln()
	fmt.Println("done")
}
