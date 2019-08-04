package main

func main() {
	nextNumber0 := getSquence()
	println(nextNumber0())
	println(nextNumber0())
	println(nextNumber0())

	nextNumber1 := getSquence()
	println(nextNumber1())
	println(nextNumber1())
}

func getSquence() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
