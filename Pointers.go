package main

import "fmt"

func transByVal(i int) {
	i = 0
}
func transByPointer(ptr *int) {
	*ptr = 0
}

func main() {

	i:=1
	fmt.Println("i:",i)
	transByVal(i)
	fmt.Println("i:",i)
	transByPointer(&i)
	fmt.Println("i:",i)

	fmt.Println("&i:",&i)


}
