package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		// wg.Add(1)
		for i := 0; i < 10; i++ {
			ch <- i
		}
		wg.Done()
	}()
	go func() {
		// wg.Add(1)
		for j := 0; j < 10; j++ {
			ch <- j
		}
		wg.Done()
	}()
	go func() {
		wg.Wait()
		close(ch)
	}()
	for n := range ch {
		fmt.Println("Data :", n)
	}

}
