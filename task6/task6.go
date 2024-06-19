package main

import (
	"context"
	"fmt"
	"sync"
)

func main() {
	//We just wait for it to end
	var wg sync.WaitGroup
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		fmt.Println("N1: I'll stop by myself")
	}(&wg)

	//Use channel
	stopCh := make(chan bool, 1)
	wg.Add(1)
	go func(wg *sync.WaitGroup, stopCh chan bool) {
		defer wg.Done()
		fmt.Print("N2: Waiting for channel")
		for {
			select {
			case <-stopCh:
				fmt.Println()
				return
			default:
				fmt.Print(".")
			}
		}
	}(&wg, stopCh)
	stopCh <- true

	//Cancel context
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go func(ctx context.Context, wg *sync.WaitGroup) {
		defer wg.Done()
		fmt.Println("N3: Waiting for cancel")
		for {
			select {
			case <-ctx.Done():
				return
			default:
				fmt.Print(".")
			}
		}
	}(ctx, &wg)
	cancel()
	wg.Wait()
}
