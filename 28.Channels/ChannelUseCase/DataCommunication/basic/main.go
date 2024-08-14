/*
Channels enable Go routines to send and receive data with each other.
You can think of a channel as a pipe where one Go routine sends data into the pipe,
and another Go routine receives it from the other end.
*/
package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		ch <- 108 //send data to the channel
	}()
	res := <-ch //recived data to the channel
	fmt.Println("channel value ", res)
}
