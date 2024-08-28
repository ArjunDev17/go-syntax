/*
Input: arr[] = [3, 4, 7, 1, 2, 9, 8]
Output: true
Explanation: (3, 7) and (9, 1) are the pairs whosesum are equal.
*/
package main

import "fmt"

func main() {
	fmt.Println("Sum of pair :")
	arr := []int{3, 4, 7, 1, 2, 9, 8}
	fmt.Println("ar :", arr)
	isAvailable := pairSum(arr)
	fmt.Println("isAvalable :", isAvailable)
}
func pairSum(num []int) bool {
	fmt.Println("num :", num)
	for i := 0; i < len(num); i++ {
		t := num[i] + num[i+1]
		for j := i + 1; j < len(num); j++ {
			t1 := num[j] + num[j+1]
			if t == t1 {
				return true
			}
		}
	}
	return false
}
