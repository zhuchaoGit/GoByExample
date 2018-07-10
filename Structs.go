package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Alice", age: 20})
	fmt.Println(person{name: "Fred"})
	fmt.Println(&person{name: "Ann"})
	s := person{name: "Sean", age: 20}
	fmt.Println(&s)
	fmt.Println(s.name)
	pointer := &s
	fmt.Println(pointer.name)
	pointer.age = 50
	fmt.Println(*pointer)
}
