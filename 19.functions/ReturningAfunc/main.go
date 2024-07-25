package main

import "fmt"

func main() {
	x := sum()
	fmt.Println("x :", x)
	y := sub()
	fmt.Println("y: ", y())
	fmt.Printf("%T :", y)
}
func sum() int {
	return 5 + 10
}
func sub() func() int {
	return func() int {
		return 10 - 5
	}
}
