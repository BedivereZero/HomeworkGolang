package main

func main() {
	var c1, c2, c3 chan int
	var i1, i2 int
	select {
		case i1 = <- c1:
			println("recceived", i1, "from c1")
		case c2 <- i2:
			println("sent", i2, "to c2")
		case i3, ok := (<- c3):
			if ok {
				println("received", i3, "from c3")
			} else {
				println("c3 is closed")
			}
		default:
			println("no communication")
	}
}
