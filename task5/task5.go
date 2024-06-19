package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func worker(ctx context.Context, wg *sync.WaitGroup, c chan int) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker stopped")
			return
		case msg := <-c:
			fmt.Println(msg)
		}
	}
}

func main() {
	var wg sync.WaitGroup

	// After 10 seconds context is Done by timeout.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	/*
	 * cancel() is deferred to release recources
	 * if worker completes before timeout elapses.
	 * In our case it will never happen
	 */
	defer cancel()

	ch := make(chan int)
	wg.Add(1)
	go worker(ctx, &wg, ch)
	for i := 10; i > 0; i-- {
		ch <- i
		time.Sleep(time.Second)
	}
	wg.Wait()
	fmt.Println("Bye!")
}
