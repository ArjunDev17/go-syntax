package main

import "fmt"

func main() {
	fmt.Println("Factorial")
	res := Factorial(4)
	fmt.Println("Factorial of this number is :", res)
}
func Factorial(number int) (res int) {
	var produnct int = 1
	for i := number; i > 0; i-- {
		produnct *= i
	}
	return produnct
}
