package main
import (
    "fmt"
    "context"
    "os"
    "os/signal"
    "syscall"
    "sync"
)

func worker(ctx context.Context, wg *sync.WaitGroup, c chan string) {
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
    //For cancelling on Ctrl+C
    ctx, cancel := context.WithCancel(context.Background()) 

    //Channel to notify us about SIGINT(Ctrl+C)
    sigCh := make(chan os.Signal, 1)
    signal.Notify(sigCh, syscall.SIGINT)

    var n int
    fmt.Println("Enter amount of workers:")
    fmt.Scanln(&n)
    ch := make(chan string, 1)

    for i:=0; i<n; i++ {
	wg.Add(1)
	go worker(ctx, &wg, ch)
    }

    //Read input until stopped
    wg.Add(1)
    go func(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	var s string
	for {
	    select {
	    case <-ctx.Done():
		return
	    default:
		fmt.Scanln(&s)
		ch <- s
	    }
	}
    }(ctx, &wg)

    //Block until SIGINT
    sig := <-sigCh
    fmt.Println(sig)
    //Clearup
    cancel()
    wg.Wait()
}
