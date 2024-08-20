package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		ch <- 108
	}()
	fmt.Println("ch  :", <-ch)
}
