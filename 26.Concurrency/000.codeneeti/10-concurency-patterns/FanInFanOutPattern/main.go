package main

import (
	"fmt"
	"sync"
)

// generator conver a list of integers to channel
func generator(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, v := range nums {
			out <- v
		}

		close(out)
	}()
	return out
}
func square(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for value := range in {
			out <- value * value
		}
		close(out)
	}()
	return out
}
func merge(cs ...<-chan int) <-chan int {

	//impliment fan-in
	//merge list of cahnnel into single channel
	out := make(chan int)

	var wg sync.WaitGroup
	output := func(c <-chan int) {
		for n := range c {
			out <- n
		}
		wg.Done()
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
	in := generator(2, 3)
	in1 := generator(4, 5)

	//TODO fan out square stage to run two instances
	ch1 := square(in)
	ch2 := square(in1)

	// fan in results of square
	mergeOut := merge(ch1, ch2)
	for value := range mergeOut {
		fmt.Println("value :", value)
	}
}
