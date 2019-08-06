package main

import "fmt"

func main() {
	nums := [] int {0, 2, 4, 6, 8}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	println("Sum:", sum)

	for i, num := range nums {
		if ( num == 4 ) {
			println("Index:", i)
		}
	}

	kvs := map[string] string {"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	for i, c := range "go" {
		println(i, c)
	}
}
