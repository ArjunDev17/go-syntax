package main

import "fmt"

func main() {
	fmt.Println("Fan in Fan out")
	// res := generate(2, 3)
}
func generate(recive ...int) chan int {
	out := make(chan int)
	for va := range recive {
		out <- va
	}
	close(out)
	return out
}
