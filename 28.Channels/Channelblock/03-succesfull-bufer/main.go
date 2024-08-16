package main

import "fmt"

func main() {
	ch := make(chan int, 1)
	ch <- 108
	fmt.Println("result ", <-ch)
}
