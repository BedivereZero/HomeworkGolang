package main

import "fmt"

func printSlice(x [] int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

func main() {
	var numbers0 [] int = make([] int, 3, 5)
	printSlice( numbers0 )

	var numbers1 [] int
	printSlice( numbers1 )
	if (numbers1 == nil) {
		fmt.Printf("Slice is nil\n")
	}

	var numbers2 [] int = [] int {0, 1, 2, 3, 4, 5, 6, 7}
	printSlice( numbers2 )

	/* print original slice */
	fmt.Println("numbers2:\t", numbers2)

	/* print slice[1:4] */
	fmt.Println("numbers2[1:4]:\t", numbers2[1:4])

	/* print slice[:3] */
	fmt.Println("numbers2[:3]:\t", numbers2[:3])

	/* print slice[4:] */
	fmt.Println("numbers2[4:]:\t", numbers2[4:])

	var numbers3 [] int = make([] int, 0, 5)
	printSlice( numbers3 )

	numbers4 := numbers2[:2]
	printSlice( numbers4 )

	numbers5 := numbers2[2:5]
	printSlice( numbers5 )

	var numbers6 [] int
	printSlice( numbers6 )

	/* append to nil slice */
	numbers6 = append( numbers6, 0 )
	printSlice( numbers6 )

	/* append one element */
	numbers6 = append( numbers6, 1 )
	printSlice( numbers6 )

	/* append multi elements */
	numbers6 = append( numbers6, 2, 3, 4 )
	printSlice( numbers6 )

	/* create slice numbers7 two cap of numbers6 */
	numbers7 := make( [] int, len( numbers6 ), ( cap( numbers6 ) ) * 2 )

	/* copy elements from numbers6 to numbers7 */
	copy( numbers7, numbers6 )
	printSlice( numbers7 )
}
