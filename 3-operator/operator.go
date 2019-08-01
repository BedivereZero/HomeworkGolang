package main

import "fmt"

func main() {

	var a int = 21
	var b int = 10
	var c int

	c = a + b
	fmt.Printf("0 - c: %d\n", c)
	c = a - b
	fmt.Printf("1 - c: %d\n", c)
	c = a * b
	fmt.Printf("2 - c: %d\n", c)
	c = a / b
	fmt.Printf("3 - c: %d\n", c)
	c = a % b
	fmt.Printf("4 - c: %d\n", c)
	a++
	fmt.Printf("5 - c: %d\n", a)
	a = 21
	a--
	fmt.Printf("6 - c: %d\n", a)

	a, b = 21, 10
	fmt.Println("a, b:", a, b)
	if(a==b) {
		fmt.Printf("0 - a equals b\n")
	} else {
		fmt.Printf("0 - a is not equal to b\n")
	}
	if ( a < b ) {
		fmt.Printf("1 - a is less than b\n")
	} else {
		fmt.Printf("1 - a is not less than b\n")
	}
	if ( a > b ) {
		fmt.Printf("2 - a is greater than b\n")
	} else {
		fmt.Printf("2 - a is not greater than b\n")
	}
	a, b = 5, 20
	fmt.Println("a, b:", a, b)
	if ( a <= b ) {
		fmt.Printf("3 - a is less than or equal to b\n")
	}
	if ( b >= a ) {
		fmt.Printf("4 - b is greater than or equal to a\n")
	}
}
