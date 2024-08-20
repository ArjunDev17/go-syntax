package main

import "fmt"

func sum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}
func main() {
	fmt.Println("sum is :", sum(2, 3))
	fmt.Println("sum is :", sum(21, 3))
	fmt.Println("sum is :", sum(2, 13))
	fmt.Println("sum is :", sum(12, 3))
	fmt.Println("sum is :", sum(23, 1))
}
