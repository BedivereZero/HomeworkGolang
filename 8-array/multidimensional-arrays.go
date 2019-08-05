package main

import "fmt"

func main() {
	var a [5][4] int

	for i := 0; i < 5; i++ {
		for j := 0; j < 4; j++ {
			a[i][j] = i + j
		}
	}

	for i := 0; i < 5; i++ {
		for j := 0; j < 4; j++ {
			if j > 0 {
				fmt.Printf("\t")
			}
			print(a[i][j])
		}
		fmt.Printf("\n")
	}
}
