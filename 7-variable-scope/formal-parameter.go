package main

import "fmt"

var a int = 20

func sum(a, b int) int {
	fmt.Printf("a in sum:\t%d\n", a)
	fmt.Printf("b in sum:\t%d\n", b)
	return a + b
}

func main() {
	var a, b, c int = 10, 20, 0
	fmt.Printf("a in main:\t%d\n", a)
	c = sum(a, b)
	fmt.Printf("c in main:\t%d\n", c)
}
