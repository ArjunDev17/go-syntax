package main

import (
	"fmt"
	"runtime"
	"time"

	"sync"
)

// generator conver a list of integers to channel
func generator(done <-chan struct{}, nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, v := range nums {
			select {
			case out <- v:
			case <-done:
				return
			}
		}
	}()
	return out
}
func square(done <-chan struct{}, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for value := range in {
			select {
			case out <- value * value:
			case <-done:
				return
			}
		}

	}()
	return out
}
func merge(done <-chan struct{}, cs ...<-chan int) <-chan int {

	//impliment fan-in
	//merge list of cahnnel into single channel

	out := make(chan int)

	var wg sync.WaitGroup
	output := func(c <-chan int) {
		defer wg.Done()
		for n := range c {
			select {
			case out <- n:
			case <-done:
				return

			}

		}

	}
	go func() {
		wg.Wait()
		close(out)
	}()
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}
	return out
}
func main() {
	done := make(chan struct{})
	in := generator(done, 2, 3)
	in1 := generator(done, 4, 5)

	//TODO fan out square stage to run two instances
	ch1 := square(done, in)
	ch2 := square(done, in1)

	// fan in results of square
	mergeOut := merge(done, ch1, ch2)
	for value := range mergeOut {
		fmt.Println("value :", value)
	}
	close(done)
	time.Sleep(1 * time.Millisecond)
	fmt.Println("r :", runtime.NumGoroutine())
}
