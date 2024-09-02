package main

import "fmt"

func linearSearch(arr []int, target int) int {
	for i, v := range arr {
		if v == target {
			return i
		}
	}
	return 0
}
func main() {
	// arr:=[8]int{2,4,5,11,1}//cannot use arr (variable of type [8]int) as []int
	// value in argument to linearSearchcompilerIncompatibleAssign
	arr := []int{2, 4, 5, 11, 1}
	index := linearSearch(arr, 2)
	fmt.Println("Value find at index :", index)
}
