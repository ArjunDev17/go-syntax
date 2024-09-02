package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	go bar()
	go foo()
	wg.Wait()
}
func bar() {
	fmt.Println("inside bar")
	for i := 0; i < 10; i++ {
		fmt.Println("value of i:", i)
	}
	wg.Done()
}
func foo() {
	fmt.Println("inside foo")
	for i := 0; i < 10; i++ {
		fmt.Println("value of J:", i)
	}
	wg.Done()
}
