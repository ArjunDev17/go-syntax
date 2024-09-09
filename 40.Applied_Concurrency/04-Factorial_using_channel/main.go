package main

import "fmt"

func main() {
	fmt.Println("Factorial using channel")
	res := Factorial(4)
	for value := range res {
		fmt.Println("Factorial of this number is :", value)
	}

}
func Factorial(number int) chan int {
	out := make(chan int)
	var produnct int = 1
	go func() {
		for i := number; i > 0; i-- {
			produnct *= i
		}
		out <- produnct
		close(out)
	}()

	return out
}
