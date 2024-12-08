package main

import "fmt"

func main() {

	ch := make(chan int)
	go func(a, b int) {
		c := a + b
		ch <- c
	}(2, 3)
	fmt.Printf("computed value :%v\n", <-ch)
}
