package main

import "fmt"

func main() {
	ch := make(chan<- int)
	ch <- 108
	res := <-ch //cannot receive

	fmt.Println("ch cant' recive ", <-ch)
}
