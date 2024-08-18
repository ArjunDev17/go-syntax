package main

import "fmt"

func main() {
	ch := make(<-chan int) //cannot send to receive-only channel

	ch <- 108 //cannot send to receive-only channel

	res := <-ch //cannot receive

	fmt.Println("ch cant' recive ", <-ch)
}
