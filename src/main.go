package main

import "fmt"

func main() {
	// Statement of constants
	const pi float64 = 3.14
	const pi2 = 3.14

	fmt.Println("Pi", pi)
	fmt.Println("Pi2", pi2)

	// Declaration of integer constants
	base := 12
	var height int = 14
	var area int

	fmt.Println(base, height, area)

	// Zero values
	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	// Square Area
	const baseSquare = 10
	squareArea := baseSquare * baseSquare
	fmt.Println("Square area:", squareArea)

}
