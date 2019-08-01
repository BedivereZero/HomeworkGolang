package main

import "fmt"

func main() {
	var a, b bool = true, false
	if ( a && b ) {
		fmt.Printf("0 - condition is true\n")
	}
	if ( a || b ) {
		fmt.Printf("1 - condition is true\n")
	}
	a, b = false, true
	if ( a && b ) {
		fmt.Printf("2 - condition is true\n")
	} else {
		fmt.Printf("2 - condition is false\n")
	}
	if ( !( a && b ) ) {
		fmt.Printf("3 - condition is true\n")
	}
}
