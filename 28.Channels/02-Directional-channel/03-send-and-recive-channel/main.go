package main

import "fmt"

func main() {
	//bidirectional channel
	c := make(chan int)

	//send only
	cs := (chan<- int)(c)

	//recive only
	cr := (<-chan int)(c)

	//send data into channel
	go func() {
		cs <- 108
	}()
	recivedData := <-cr
	fmt.Println("data is  :", recivedData)

}
