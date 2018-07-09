package main

import "fmt"

func fact(n int) int {
	if n <= 0 {
		return 1
	}
	//recursion
	fmt.Println(n * fact(n-1))
	//double recursion
	return n * fact(n-1)
}

func main() {
	fact(5)
}
