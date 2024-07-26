package main

import "fmt"

type myNum interface {
	int | float64
}

func genericsCal[T myNum](a, b T) T {
	return a - b
}
func main() {

	fmt.Println("--------------------")
	fmt.Println("by generics :", genericsCal(4, 2))
	fmt.Println("by generics :", genericsCal(4.2, 2.1))
}
