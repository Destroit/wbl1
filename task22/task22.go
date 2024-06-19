package main

import (
	"log"
	"math/big"
)

func main() {
	// Operation with big numbers are done with math/big
	log.SetFlags(0)
	x, ok := big.NewInt(0).SetString("999999999999999999", 10)
	if !ok {
		log.Fatal("error setting x")
	}
	y, ok := big.NewInt(0).SetString("111111111111111111", 10)
	if !ok {
		log.Fatal("error setting y")
	}

	log.Println("Sum: ", big.NewInt(0).Add(x, y))
	log.Println("Sub: ", big.NewInt(0).Sub(x, y))
	log.Println("Mul: ", big.NewInt(0).Mul(x, y))
	log.Println("Div: ", big.NewInt(0).Div(x, y))
}
