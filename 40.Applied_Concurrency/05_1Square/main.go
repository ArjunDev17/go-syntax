package main

import "fmt"

func main() {
	for value := range square(gen(2, 3)) {
		fmt.Println("Square is :", value)
	}
}
func gen(number ...int) chan int {
	out := make(chan int)
	go func() {
		for _, part := range number {
			out <- part
		}
		close(out)
	}()
	return out
}
func square(in chan int) chan int {
	res := make(chan int)
	go func() {
		for value := range in {
			res <- value * value
		}
		close(res)
	}()
	return res
}
