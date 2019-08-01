package main

import "fmt"
import "unsafe"

const (
	d = "a数字常量可以用"
	e = len(d)
	f = unsafe.Sizeof(d)
)

const (
	g = iota
	h = iota
	i = iota
)

const (
	j = iota
	k
	l
)

const (
	m = iota
	n
	o = "ha"
	p
	q = 100
	r = iota
	s
)

func main() {
	println("constant")
	const LENGTH int = 10
	const WEIDTH int = 20
	var area int
	const a, b, c = 1, false, "string"
	area = LENGTH * WEIDTH
	fmt.Printf("Area: %d\n", area)
	println(a, b, c)

	println("enum")
	println(d, e, f)

	println("iota")
	println(g, h, i)
	println(j, k, l)
	println(m, n, o, p, q, r, s)
}
