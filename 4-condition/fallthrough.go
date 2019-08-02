package main

func main() {
	switch{
		case false:
			println("0 - case condition is false")
			fallthrough
		case true:
			println("1 - case condition is true")
			fallthrough
		case false:
			println("2 - case condition is false")
			fallthrough
		case true:
			println("3 - case condition is true")
		case false:
			println("4 - case condition is false")
		default:
			println("5 - default case")
	}
}
