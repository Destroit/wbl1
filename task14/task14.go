package main

import (
	"fmt"
)

func getType(i interface{}) {
	// this switch is a set of type assertions
	switch v := i.(type) {
	case int:
		fmt.Println("Is int: ", v)
	case string:
		fmt.Println("Is string: ", v)
	case bool:
		fmt.Println("Is bool: ", v)
	case chan int:
		fmt.Println("Is chan int: ", v)
	default:
		fmt.Printf("Unknown type %T\n", v)
	}
}
func main() {
	getType(1)
	getType("hello")
	getType(true)
	v := make(chan int)
	getType(v)
}
