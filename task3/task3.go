package main

import (
	"fmt"
)

func square(x int, c chan int) {
	c <- x * x
}

func main() {
	sum := 0
	num := []int{2, 4, 6, 8, 10}

	// Creating buffered channel for squares. But unbuffered works as well
	ch := make(chan int, len(num))
	for _, v := range num {
		// Add 1 to count of goroutines to wait for
		go square(v, ch)
	}
	for range num {
		// Blocks until channel recieves value
		sum += <-ch
	}
	fmt.Println(sum)
}
