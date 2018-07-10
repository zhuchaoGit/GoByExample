package main

import (
	"errors"
	"fmt"
)

type bizError struct {
	code    int
	message string
}

func (be *bizError) Error() string {
	return fmt.Sprintf("%d-%s", be.code, be.message)
}

func f1(arg int) (int, error) {
	if arg > 40 {
		return arg, errors.New("beyond range")
	}
	return arg, nil
}
func f2(arg int) (int, error) {
	if arg > 40 {
		return arg, &bizError{arg, "biz error"}
	}
	return arg, nil
}

func run(f func(arg int) (int, error)) {
	sl := []int{7, 42}
	for _, i := range sl {
		if r, e := f(i); e != nil {
			fmt.Println("error:", r)
		} else {
			fmt.Println("work:", r)
		}
	}
}
func main() {
	run(f1)
	run(f2)
	_, e := f2(41)
	if be, ok := e.(*bizError); ok {
		fmt.Printf("code:%d,message:%s", be.code, be.message)
	}
}
