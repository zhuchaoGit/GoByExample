package main

import (
	"fmt"
	"reflect"
)

type GS interface {
	Get() int
	Set(int)
	GetDouble() int
}

type G interface {
	Get() int
	GetDouble() int
}

type S struct {
	val int
}

func (s *S) Get() int {
	return s.val
}
func (s *S) Set(i int) {
	s.val = i
}
func (s *S) GetDouble() int {
	return s.val * 2
}

func fgs(gs GS) {
	gs.Set(1)
	fmt.Println("gs:", reflect.TypeOf(gs).String(), gs.Get())
	fmt.Println("double gs:", reflect.TypeOf(gs).String(), gs.GetDouble())
}

func fg(g G) {
	fmt.Println("g:", reflect.TypeOf(g).String(), g.Get())
	fmt.Println("double g:", reflect.TypeOf(g).String(), g.GetDouble())
}

func main() {
	s := S{2}
	fg(&s)
	fgs(&s)
}
