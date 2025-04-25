package main

import "fmt"

func Add(a, b int) int {
	return a + b
}
func main() {
	result := Add(2, 4)
	fmt.Println("addition of two number is :", result)
}
