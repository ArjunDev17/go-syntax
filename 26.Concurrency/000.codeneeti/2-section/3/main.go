package main

import (
	"fmt"
	"sync"
)

/*
Todo:modify the program
to print data value is one
deterministically
*/
func main() {
	var data int
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		data++
	}()
	// if data == 0 {
	// 	fmt.Printf("data is :%d\n", data)
	// }
	wg.Wait()
	fmt.Printf("data is :%d\n", data)
	fmt.Println("Done....")
}
