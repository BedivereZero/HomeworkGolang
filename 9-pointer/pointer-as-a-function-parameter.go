package main

func swap(x, y * int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}

func main() {
	var a, b int = 100, 200
	println("Before:\t", a, b)
	swap(&a, &b)
	println("After:\t", a, b)
}
