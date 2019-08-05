package main

import "fmt"

var g int
var h int = 20

func main() {
	var a, b int
	a = 10
	b = 20
	g = a + b
	var h int = 30
	fmt.Printf("Result: a = %d, b = %d and c = %d\n", a, b, g)
	fmt.Printf("Result: h = %d\n", h)
}
