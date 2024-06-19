package main

import (
	"fmt"
	"sync"
)

// Put squares of values from in channel, to out channel
func square(wg *sync.WaitGroup, in chan int, out chan int) {
	defer wg.Done()
	for {
		x, ok := <-in
		if !ok {
			//in is closed, no more values
			close(out)
			return
		}
		out <- x * x
	}
}

func printer(wg *sync.WaitGroup, out chan int) {
	defer wg.Done()
	for {
		x, ok := <-out
		if !ok {
			// out is closed
			return
		}
		fmt.Println(x)
	}
}

func main() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	var wg sync.WaitGroup

	inCh := make(chan int, 1)
	outCh := make(chan int, 1)

	wg.Add(2)
	go square(&wg, inCh, outCh)
	go printer(&wg, outCh)

	for _, v := range arr {
		inCh <- v
	}
	close(inCh)
	wg.Wait()
}
