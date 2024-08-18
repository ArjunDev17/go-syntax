// given an array . you have to find the sum of range mentioned in input.

// [4:27 PM] Prachi Pandey
// input - [1,2,3,4,5,6,7]
// range - [2,4]
// sum - 9

package main

import "fmt"

func main() {

	arr := []int{1, 2, 3, 4, 5, 6, 7}
	start := 2
	end := 4
	result := sumOfRange(arr, start, end)
	fmt.Printf("The sum of elements from given range :%d", result)
}

func sumOfRange(arr []int, start, end int) int {
	sum := 0
	for i := start; i <= end; i++ {
		sum += arr[i]
	}
	return sum
}
