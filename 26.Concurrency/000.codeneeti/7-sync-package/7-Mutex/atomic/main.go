package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	// "sync/atomic"
)

func main() {
	runtime.GOMAXPROCS(4)
	var count uint32
	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				atomic.AddUint32(&count, 1)
				// count++

			}
		}()

	}
	wg.Wait()
	fmt.Printf("count: %v\n", count)
}
