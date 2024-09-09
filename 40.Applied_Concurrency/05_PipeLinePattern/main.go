package main

import "fmt"

func main() {

	sc := gen(2, 3, 4, 6)
	squar := squar(sc)
	for value := range squar {
		fmt.Println("squar of this number :", value)
	}
}
func gen(number ...int) chan int {
	out := make(chan int)
	go func() {
		for _, n := range number {
			out <- n
		}
		close(out)
	}()
	return out
}
func squar(in chan int) chan int {
	ch := make(chan int)
	go func() {
		for value := range in {
			ch <- value * value
		}
		close(ch)
	}()
	return ch
}
