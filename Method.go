package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.height * r.width
}
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{20, 30}
	fmt.Println("area", r.area())
	rsp := &r
	fmt.Println("perm", rsp.perim())
}
