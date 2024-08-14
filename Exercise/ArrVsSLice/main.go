package main

import "fmt"

func main() {
	var myArray [3]int // This defines an array of 3 integers
	myArray[0] = 10
	myArray[1] = 20
	myArray[2] = 30
	// myArray[3] = 30 invalid argument: index 3 out of bounds

	fmt.Println("Array:", myArray) // Output: Array: [10 20 30]
}
