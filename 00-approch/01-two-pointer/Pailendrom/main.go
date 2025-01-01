package main

import "fmt"

func isPalendrom(str string) bool {
	left := 0
	right := len(str) - 1

	for left < right {

		if str[left] != str[right] {

			return false
		}
		left++
		right--
	}
	return true
}
func main() {
	fact := isPalendrom("madamwe")
	fmt.Println("is palendrom :", fact)
}
