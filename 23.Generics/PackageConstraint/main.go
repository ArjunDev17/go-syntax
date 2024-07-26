package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// Define a generic function to perform subtraction.
func genericsCal[T myNum](a, b T) T {
	return a - b
}

// Define an interface to accept either integer or floating-point types.
type myNum interface {
	constraints.Integer | constraints.Float
}

// Define a custom type 'myAge' as an integer.
type myAge int

func main() {
	// Use the custom type 'myAge' and perform operations using generics.
	var age myAge = 25
	fmt.Println("by generics (myAge):", genericsCal(age, 2))
	fmt.Println("by generics (float64):", genericsCal(4.2, 2.1))
}
