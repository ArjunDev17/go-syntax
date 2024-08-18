package main

import "fmt"

func main() {
	// integer array 1,0,3,0,7,0,2,0
	// 1 3 7 2 0 0 0
	arr := []int{1, 0, 3, 0, 7, 0, 2, 0}
	fmt.Println("Original array :", arr)
	result := moveZeroToEnd(arr)
	fmt.Println("Array after moving zero to right side", result)
}

func moveZeroToEnd(arr []int) []int {
	nonZeroIndex := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			arr[nonZeroIndex] = arr[i]
			nonZeroIndex++
		}
	}

	for i := nonZeroIndex; i < len(arr); i++ {
		arr[i] = 0
	}
	return arr
}

// has context menu
