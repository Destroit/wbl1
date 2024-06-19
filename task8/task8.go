package main

import (
	"fmt"
	"log"
	"strings"
)

func main() {
	log.SetFlags(0)
	var (
		n     int64
		shift int
		value int
	)

	// Get number
	fmt.Print("Number: ")
	_, err := fmt.Scanln(&n)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%064b\n", n)

	// Select position (0 - right, 63 - left)
	fmt.Print("Position [0-63]: ")
	_, err = fmt.Scanln(&shift)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%064b\n", n)

	if shift > 63 || shift < 0 {
		log.Fatal("Wrong position")
	}
	// Print cursor underneath selected bit
	cursor := strings.Repeat(" ", 63-shift) + "^"
	fmt.Println(cursor)

	// Select value
	fmt.Print("Value [0/1]: ")
	_, err = fmt.Scanln(&value)
	if err != nil {
		log.Fatal(err)
	}

	/*
	 * We use selected value of bit
	 * If we want to invert bit, XOR
	 * is enough
	 */
	if value == 1 {
		n = n | (1 << shift)
	} else if value == 0 {
		n = n & (^(1 << shift))
	} else {
		log.Fatal("Wrong value")
	}
	// Print result
	fmt.Printf("%064b\n", n)
}
