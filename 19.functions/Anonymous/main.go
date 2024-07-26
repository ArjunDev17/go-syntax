package main

import "fmt"

func main() {
	foo()
	func() {
		fmt.Println("Anonymous function")
	}()
}
func foo() {
	fmt.Println("Normal function")
}
