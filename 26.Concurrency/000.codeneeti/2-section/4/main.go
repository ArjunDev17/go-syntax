package main

import (
	"fmt"
	"sync"
)

func increment(wg *sync.WaitGroup) {
	var count int // Local variable inside the function

	wg.Add(1)
	go func() {
		defer wg.Done()
		count++ // Increment the local variable
		fmt.Println("Count:", count)
	}()

	fmt.Println("Function returned")
}

func main() {
	var wg sync.WaitGroup
	increment(&wg) // Call the function
	wg.Wait()      // Wait for the Goroutine to finish
	fmt.Println("Program complete")
}
