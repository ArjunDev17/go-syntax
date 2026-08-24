// given an array . you have to find the sum of range mentioned in input.

// [4:27 PM] PP
// input - [1,2,3,4,5,6,7]
// range - [2,4]
// sum - 9

// arr := []int{1, 2, 3, 4, 5, 6, 7,8,9}  i need to perform sum of array
// 3 go ro

package main

import (
	"fmt"
	"sync"
)

func main() {
	// range - [2,4], [3,5], [4,5]
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	n := len(arr)
	partSize := n / 3

	var sum1, sum2, sum3 int
	var wg sync.WaitGroup

	wg.Add(3)

	go sumArray(arr, &wg, &sum1, 0, partSize)
	go sumArray(arr, &wg, &sum2, partSize, 2*partSize)
	go sumArray(arr, &wg, &sum3, 2*partSize, n)

	wg.Wait()

	finalSum := sum1 + sum2 + sum3
	fmt.Println(finalSum)
}

func sumArray(arr []int, wg *sync.WaitGroup, result *int, start int, end int) {
	defer wg.Done()
	for i := start; i < end; i++ {
		*result += arr[i]
	}

}
