package main

func main() {
	var a int = 10
	for a < 20 {
		println("a is", a)
		a++
		if a > 15 {
			break
		}
	}
}
