package main

import "fmt"

var (
	a, b int
	res  float64
)

func main() {
	fmt.Println("calculator")
	inputTwoNum()
	fmt.Println("your a nd b value is ", a, b)
	inputTwoNum1()
	fmt.Println("your a nd b value is ", a, b)
}
func inputTwoNum() {
	fmt.Println("enter a value :")
	fmt.Scanf("%d %d", &a, &b)
}
func inputTwoNum1() {
	fmt.Println("enter a value :")
	fmt.Scanln(&a)
	// fmt.Println("")
	fmt.Scanln(&b)
}
