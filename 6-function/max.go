package main

func main() {
	var a, b int = 100, 200
	var ret int
	ret = max(a, b)
	println("max is:", ret)
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
