package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	// Mutex for preventing simultaneous access
	mu  sync.Mutex
	cnt int64
}

func newCounter() Counter {
	return Counter{cnt: 0}
}

func adder(wg *sync.WaitGroup, c *Counter) {
	// Blocks if other goroutine used Lock earlier
	c.mu.Lock()
	defer c.mu.Unlock()
	defer wg.Done()
	c.cnt++
}

func main() {
	c := newCounter()
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go adder(&wg, &c)
	}
	wg.Wait()
	fmt.Println(c.cnt)
}
