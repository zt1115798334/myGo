package main

import "fmt"

const (
	a = "abc"
	b = true
	c = 2
)

const (
	xiao1 = iota
	xiao2 = iota
	xiao3 = iota
)

const (
	d1 = 1 << iota
	d2 = 2 << iota
	d3 = 3 << iota
	d4
)

func main() {
	fmt.Println(a, "111", b, c)
	fmt.Println(xiao1, xiao2, xiao3)
	fmt.Println(d1, d2, d3, d4)


}
