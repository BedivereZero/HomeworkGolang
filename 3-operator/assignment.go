package main

import "fmt"

func main() {
	
	var a int = 21
	var c int

	c = a
	fmt.Printf("0: %d\n", c)

	c += a
	fmt.Printf("1: %d\n", c)

	c -= a
	fmt.Printf("2: %d\n", c)

	c /= a
	fmt.Printf("3: %d\n", c)

	c = 200
	fmt.Printf("4: %d\n", c)

	c <<= 2
	fmt.Printf("5: %d\n", c)

	c >>= 2
	fmt.Printf("6: %d\n", c)

	c &= 2
	fmt.Printf("7: %d\n", c)

	c ^= 2
	fmt.Printf("8: %d\n", c)

	c |= 2
	fmt.Printf("9: %d\n", c)
}
