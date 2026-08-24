// given an array . you have to find the sum of range mentioned in input.

// [4:27 PM] Prachi Pandey
// input - [1,2,3,4,5,6,7]
// range - [2,4]
// sum - 9

// arr := []int{1, 2, 3, 4, 5, 6, 7,8,9}  i need to perform sum of array
// 3 go ro

package main

import "fmt"

func main() {
	// range - [2,4], [3,5], [4,5]
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	ranges := [][]int{
		{2, 4},
		{3, 5},
		{4, 5},
	}
	for _, r := range ranges {
		start, end := r[0], r[1]
		result := sumOfRange(arr, start, end)
		fmt.Printf("The sum of elements from given range :%d\n", result)

	}
}

func sumOfRange(arr []int, start, end int) int {
	sum := 0
	for i := start; i <= end; i++ {
		sum += arr[i]
	}
	return sum
}
