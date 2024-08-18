/*
Given an integer array nums, move all 0's to the end
of it while maintaining the relative order of the non-zero elements.
Note that you must do this in-place without making a copy of the array.

Example 1:
Input: nums = [0,1,0,3,12]
Output: [1,3,12,0,0]
*/
package main

import (
	"fmt"
)

func main() {
	// nums := []int{1, 4, 2, 0, 7, 4, 3, 0, 1, 1}
	// moveZeroToRight(nums)
	// fmt.Println(nums)
	for i := 0; i < 0; i++ {
		go printHell()
	}

}
func printHell() {
	fmt.Println("Hello world")
}
func moveZeroToRight(nums []int) {
	n := len(nums)
	if n == 0 {
		return
	}

	nonZeroIndex := 0

	for i := 0; i < n; i++ {
		if nums[i] != 0 {
			nums[nonZeroIndex], nums[i] = nums[i], nums[nonZeroIndex]
			nonZeroIndex++
		}
	}
}
