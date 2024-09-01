package main

import "fmt"

func main() {
	ch := incrementor()
	res := puller(ch)
	for value := range res {
		fmt.Println("value :", value)
	}
}
func incrementor() chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}
func puller(ch chan int) chan int {
	out := make(chan int)
	go func() {
		var sum int
		for value := range ch {
			sum += value
		}
		out <- sum
		close(out)
	}()
	return out
}
