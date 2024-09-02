package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}
func main() {
	wg.Add(1)
	go bar()
	go foo()
	wg.Wait()
}
func bar() {

	for i := 0; i < 10; i++ {
		fmt.Println("value of i:", i)
		time.Sleep(time.Duration(3 * time.Microsecond))
	}
	wg.Done()
}
func foo() {

	for i := 0; i < 10; i++ {
		fmt.Println("value of J:", i)
		time.Sleep(time.Duration(20 * time.Microsecond))
	}
	wg.Done()
}
