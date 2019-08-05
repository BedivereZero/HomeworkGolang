package main

import "fmt"

func main() {
	var a int
	var p * int
	var pp ** int

	a = 3000
	p = &a
	pp = &p

	fmt.Printf("a  = %d\n", a)
	fmt.Printf("p  = %d\n", *p)
	fmt.Printf("pp = %d\n", **pp)
}
