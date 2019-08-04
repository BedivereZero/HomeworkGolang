package main

import "math"

func main() {
	getSquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}
	println(getSquareRoot(199 * 199))
}
