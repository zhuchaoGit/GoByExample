package main

import (
	"fmt"
	"reflect"
)

type GS interface {
	Get() int
	Set(int)
}

type G interface {
	Get() int
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

func fgs(gs GS) {
	gs.Set(1)
	fmt.Println("gs:", reflect.TypeOf(gs).String(), gs.Get())
}

func fg(g G) {
	fmt.Println("g:", reflect.TypeOf(g).String(), g.Get())
}

func main() {
	s := S{2}
	fg(&s)
	fgs(&s)
}
