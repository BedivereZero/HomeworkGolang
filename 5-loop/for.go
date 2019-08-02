package main

func main() {
	var a int = 1
	var b int = 15

	println("location of a:", &a)

	numbers := [6]int{1, 2, 3, 5}

	println("---")

	for a := 0; a < 10; a++ {
		println("a:", a)
		println("location of a:", &a)
	}

	println("---")

	for a < b {
		a++
		println("a:", a)
		println("location of a:", &a)
	}

	println("---")

	for i, x := range numbers {
		println(i, ":", x)
	}
}
