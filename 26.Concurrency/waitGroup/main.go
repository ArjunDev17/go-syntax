package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("OS \t\t", runtime.GOOS)
	fmt.Println("Arch \t\t", runtime.GOARCH)
	fmt.Println("CPU----->\t\t", runtime.NumCPU())
	fmt.Println("Gorutines\t\t", runtime.NumGoroutine())
	wg.Add(1)
	go foo()
	bar()
	fmt.Println("CPU---->\t\t", runtime.NumCPU())
	fmt.Println("Gorutines\t\t", runtime.NumGoroutine())
	wg.Wait()
}
func foo() {
	count := 5
	for i := 0; i < count; i++ {
		fmt.Println("foo :", i)
	}
	wg.Done()
}
func bar() {
	count := 5
	for i := 0; i < count; i++ {
		fmt.Println("bar :", i)
	}
}
