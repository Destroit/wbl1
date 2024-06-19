package main

import (
	"fmt"
	"sync"
)

func square(i int, x int, wg *sync.WaitGroup) {
	// Decrement WaitGroup counter after func is finished
	defer wg.Done()
	fmt.Printf("Index:%d Square:%d\n", i, x*x)
}

func main() {
	// WaitGroup is used for tracking goroutines' completion
	var wg sync.WaitGroup
	num := []int{2, 4, 6, 8, 10}
	for i, v := range num {
		// Add 1 to count of goroutines to wait for
		wg.Add(1)
		go square(i, v, &wg)
	}

	// Blocks until counter becomes zero
	wg.Wait()
}
