package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	sum := 0
	for i, num := range nums {
		sum += num
		fmt.Println("i,val:", i, num)
	}
	fmt.Println("sum")
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Println("k,v:", k, v)
	}
	for k := range kvs {
		fmt.Println("k:", k)
	}
	for i, c := range "你好golang" {
		fmt.Println("i,c:", i, string(c))
	}
	for i, c := range []rune("你好golang") {
		fmt.Println("i,c:", i, string(c))
	}
}
