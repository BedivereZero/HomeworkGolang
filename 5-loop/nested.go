package main

func main() {
	for i := 2; i < 100; i++ {
		prime := true
		for j := 2; j <= (i / j); j++ {
			if ( i % j == 0) {
				prime = false
				break
			}
		}
		if prime {
			println(i, "is prime number")
		}
	}
}
