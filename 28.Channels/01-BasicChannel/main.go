package main

import "fmt"

func main() {
	fmt.Println("Channel")
	ch := make(chan int)
	ch <- 108
	fmt.Println(<-ch)
}
