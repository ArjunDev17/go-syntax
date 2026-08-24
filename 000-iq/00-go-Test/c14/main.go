package main

import (
	"fmt"
	"sync"
)

// create a program where you can demonstrate the race condtion
func demonstrateRaceCondition() int16 {
	var counter int16
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for range 100 {
			counter++
		}
	}()
	go func() {
		defer wg.Done()
		for range 100 {
			counter++
		}
	}()
	wg.Wait()
	return counter
}

func main() {
	res := demonstrateRaceCondition()
	fmt.Println("res :", res)

}
