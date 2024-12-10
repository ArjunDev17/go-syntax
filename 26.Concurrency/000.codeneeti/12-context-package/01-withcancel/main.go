package main

import (
	"context"
	"fmt"
)

func main() {
	/*
		generator genrate integers in a seprate goroutine and
		snds them to the returned channel
		the callers of gen need to cancel the goroutine once they consume 5 integers
		so that internal goroutine
		started by gen is not leaked
	*/
	genrate := func(ctx context.Context) <-chan int {
		out := make(chan int)
		n := 1
		go func() {
			defer close(out)
			for {
				select {
				case out <- n:
				case <-ctx.Done():
					return

				}
				n++
			}
		}()

		return out
	}
	//Create a context that is cancellable .
	ctx, cancel := context.WithCancel(context.Background())
	ch := genrate(ctx)
	for n := range ch {
		fmt.Println(" value is :", n)
		if n == 15 {
			cancel()
		}
	}
}
