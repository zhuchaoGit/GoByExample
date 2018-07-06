package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Println("write ", i, "as")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}
	whatami := func(i interface{}) {
		switch i.(type) {
		case bool:
			fmt.Println("bool ")
		case int:
			fmt.Println("int ")
		default:
			fmt.Println("unkonwn")
		}
	}
	whatami(true)
	whatami(1)
	whatami("string")
}
