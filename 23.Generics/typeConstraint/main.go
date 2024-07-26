package main

import "fmt"

func addI(a, b int) int {
	c := a + b
	return c
}
func addF(a, b float64) float64 {
	c := a + b
	return c
}
func genericsCal[T int | float64](a, b T) T {
	return a - b
}
func main() {
	fmt.Println(" normal :", addF(2, 4), addI(1, 3))
	fmt.Println("--------------------")
	fmt.Println("by generics :", genericsCal(4, 2))
	fmt.Println("by generics :", genericsCal(4.2, 2.1))
}
