package main

func main() {
	var a int = 10
	for a < 20 {
		if a == 15 {
			a++
			continue
		}
		println("a is", a)
		a++
	}
}
