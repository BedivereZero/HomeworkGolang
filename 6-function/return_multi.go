package main

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := "Apple", "Google"
	x, y := swap(a, b)
	println(a, b)
	println(x, y)
}
