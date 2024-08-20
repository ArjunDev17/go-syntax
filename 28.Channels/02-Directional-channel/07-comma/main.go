package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		ch <- 12
	}()
	res, ok := <-ch
	fmt.Println("result :", res, ok)
	close(ch)
	res1, ok := <-ch
	fmt.Println("result :", res1, ok)
}
