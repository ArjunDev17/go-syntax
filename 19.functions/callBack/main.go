package main

import "fmt"

func main() {
	x := doMath(2, 3, add)
	fmt.Println("x :", x)
	y := doMath(10, 2, sub)
	fmt.Println("y:", y)
	fmt.Println("-----------------------")
	f := increment()
	fmt.Println(" f:", f())
	fmt.Println(" f:", f())
	fmt.Println(" f:", f())
	fmt.Println("-----------------------")
	res := fact(4)
	fmt.Println("Factorial is :", res)

}
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact((n - 1))
}
func add(a, b int) int {
	return a + b
}
func sub(a, b int) int {
	return a - b
}
func doMath(a, b int, f func(a, b int) int) int {
	return f(a, b)
}
func increment() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
