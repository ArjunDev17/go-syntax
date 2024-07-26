package main

import "fmt"

// value semantics
func addOne(v int) int {
	return v + 1
}

// Pointer semantics
func addOneP(v *int) {
	*v += 1

}
func main() {
	fmt.Println("Pointer & value semantics defined")
	a := 1
	fmt.Println(a)
	fmt.Println(addOne(a))
	fmt.Println(addOne(a))
	fmt.Println(addOne(a))
	b := 1
	fmt.Println("-------// Pointer semantics------------------")
	fmt.Println(b)
	addOneP(&b)
	fmt.Println(b)
	addOneP(&b)
	fmt.Println(b)
	addOneP(&b)

	fmt.Println(b)
	fmt.Println(b)

}
