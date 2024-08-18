package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 108
	ch <- 109
	fmt.Println("res1 :", <-ch)
	fmt.Println("res2 :", <-ch)
}

// do not communicate by sharing memory instead share memory by communicating
