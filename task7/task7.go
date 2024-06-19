package main

import (
	"fmt"
	"sync"
)

type ConcurMap struct {
	/*
	 * Mutex prevents map from being accessed
	 * by multiple goroutines at the same time
	 */
	mu sync.Mutex
	v  map[int]int
}

func (cm *ConcurMap) Write(key int) {
	cm.v[key]++
}

func writer(wg *sync.WaitGroup, cm *ConcurMap, key int) {
	//Lock blocks other goroutines which use lock until we unlock mutex
	cm.mu.Lock()
	defer cm.mu.Unlock()
	defer wg.Done()
	cm.Write(key)
}

func main() {
	var (
		cm ConcurMap
		wg sync.WaitGroup
	)
	cm.v = make(map[int]int)
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go writer(&wg, &cm, i)
	}
	wg.Wait()
	fmt.Println(cm.v)
}
