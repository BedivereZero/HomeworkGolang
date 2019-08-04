package main

func swap(x, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}

func main() {
	var a, b int = 0, 1
	println("Before:", a, b)
	swap(&a, &b)
	println("After:", a, b)
}
