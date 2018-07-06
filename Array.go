package main

import "fmt"

type array []int

func (a array) Len() int {
	return len(a)
}
func main() {
	var a [5]int
	fmt.Println("empt:", a)
	a[4] = 100
	var arr array = make([]int, 5, 5)
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])
	fmt.Println("len:", len(a))
	fmt.Println("slice named len", arr.Len())
	var bytes = []byte("hello")
	fmt.Println(bytes)
	var b int = int(5)
	fmt.Println("init b:", b)
	var str = string(bytes)
	fmt.Println("str:", str)
}
