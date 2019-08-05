package main

import "fmt"

const MAX int = 3

func main() {
	a := [] int {10, 100, 1000}
	var p [MAX] * int
	for i := 0; i < MAX; i++ {
		fmt.Printf("a[%d] = %d\n", i, a[i])
	}
	for i := 0; i < MAX; i++ {
		p[i] = &a[i]
	}
	for i := 0; i < MAX; i++ {
		fmt.Printf("a[%d] = %d\n", i, *p[i])
	}
}
