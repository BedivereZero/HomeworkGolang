package main

type Circle struct {
	radius float64
}

func main() {
	var c1 Circle
	c1.radius = 10.0
	println("Area of circle is", c1.getArea())
}

func (c Circle) getArea() float64 {
	return 3.1415976 * c.radius * c.radius
}
