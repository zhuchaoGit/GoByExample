package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println("empty:", s)
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("set:", s)
	fmt.Println("get:", s[2])
	fmt.Println("len:", len(s))
	fmt.Println("cap:", cap(s))
	fmt.Println("address:", &s[0])

	s = append(s, "d")
	fmt.Println("address:", &s[0])
	fmt.Println("len:", len(s))
	fmt.Println("cap:", cap(s))

	s = append(s, "e", "f")
	fmt.Println("append:", s)
	fmt.Println("len:", len(s))
	fmt.Println("cap:", cap(s))
	fmt.Println("address:", &s[0])
	fmt.Println("address:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy", c)

	l := s[2:5]
	fmt.Println("sl1:", l)
	fmt.Println("address:", &l[0])
	fmt.Println("address:", &s[2])

	l = s[:5]
	fmt.Println("sl2:", l)
	l = s[2:]
	fmt.Println("sl3:", l)
	t := []string{"g", "h", "i"}
	fmt.Println("dcl", t)

	triangle := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		triangle[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			triangle[i][j] = i + j
		}
	}
	fmt.Println("triangle:", triangle)
}
