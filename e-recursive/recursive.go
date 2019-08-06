package main

func factorial(n uint64) uint64 {
	if (n == 0) {
		return 1
	} else {
		return n * factorial( n - 1 )
	}
}

func fibonacci(n uint64) uint64 {
	if (n < 2) {
		return n
	} else {
		return fibonacci( n - 1 ) + fibonacci( n - 2 )
	}
}

func main() {
	var i uint64 = 15
	println("factorial:", i, factorial(i))
	println("fibonacci:", i, fibonacci(i))
}
