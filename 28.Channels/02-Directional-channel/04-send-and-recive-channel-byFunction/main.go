package main

import "fmt"

func main() {
	c := make(chan int)
	go send(c)
	recive(c)
}
func send(c chan<- int) {
	c <- 100
}
func recive(c <-chan int) {
	fmt.Println("value is :", <-c)
}
