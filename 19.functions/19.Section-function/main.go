package main

import "fmt"

func main() {
	fmt.Println("all about function")
	vardikFunc(1, 2, 3, 4)
	num := []int{1, 2, 3, 4}
	a, b, c := num[0], num[1], num[2:]
	fmt.Println(" :", a, b, c)

	unfurLingASlice(num...)
}
func vardikFunc(x ...int) {
	sum := 0
	for _, value := range x {
		sum += value
	}
	fmt.Println("sum :", sum)
}
func unfurLingASlice(num ...int) {
	sum := 0
	for _, value := range num {
		sum += value
	}
	fmt.Println("sum :", sum)
}
