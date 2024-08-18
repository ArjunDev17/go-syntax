package main

import "fmt"

func main() {
	ch := make(chan int)
	ch <- 42
	fmt.Println("ch :", ch)
}
