package main

import "fmt"

func main() {
	// Hidden XOR
	a := 1
	b := 2
	fmt.Println("before:", a, b)
	a, b = b, a
	fmt.Println("after:", a, b)

	// Manual XOR swap
	x := 3
	y := 4
	fmt.Println("before:", x, y)
	x = y ^ x
	y = x ^ y
	x = y ^ x
	fmt.Println("after:", x, y)
}
