package main

import "fmt"

var (
	multiVarInt int
	multiVarStr bool
)

func main()  {

	var a string = "Variable"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d int
	fmt.Println(d)

	var e bool
	fmt.Println(e)

	var f int
	var g float64
	var h bool
	var i string
	fmt.Printf("%v %v %v %q\n", f, g, h, i)

	var j = true
	fmt.Println(j)

	newVarInt := 1
	newVarStr := "Variables"
	fmt.Println(newVarInt)
	fmt.Println(newVarStr)

	fmt.Println(multiVarInt)
	fmt.Println(multiVarStr)

	var valueTypeI int
	var valueTypeJ int
	valueTypeI = 1
	valueTypeJ = 2
	fmt.Println(valueTypeI, valueTypeJ)
	valueTypeI = 3
	fmt.Println(valueTypeI, valueTypeJ)
}
