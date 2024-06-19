package main

import (
	"fmt"
	"time"
)

func sleep(td time.Duration) {
	<-time.After(td)
}

func main() {
	for i := 5; i > 0; i-- {
		fmt.Println(i)
		sleep(1 * time.Second)
	}
	fmt.Println("Bye!")
}
