package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}
func plusMore(a, b, c int) int {
	return a + b + c
}
func vals() (int, int) {
	return 0, 1
}

func sum(nums ... int) int {
	sum := 0
	for num := range nums {
		sum += num
	}
	return sum
}
func main() {
	fmt.Println("1+2=", plus(1, 2))
	fmt.Println("1+2+3=", plusMore(1, 2, 3))
	a, b := vals()
	fmt.Println("vals:", a, b)

	nums := []int{1, 2, 3}
	fmt.Println("sum:", sum(nums...))
	fmt.Println("sum:", sum(1, 2))
}
