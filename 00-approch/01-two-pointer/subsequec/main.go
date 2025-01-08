package main

import (
	"fmt"
)

func isSubsequence(s string, t string) bool {
	var count int = 0
	for i := 0; i < len(s); i++ {
		for j := count; j < len(t); j++ {
			if s[i] == t[j] {
				count++
			}
		}
	}
	if count == len(s) {
		return true
	}
	return false
}
func main() {
	nums := []int{1, 2, 3, 4}
	n := len(nums) - 1
	fmt.Println(" length :", n)
	s := "abc"
	t := "ahabgdc"
	res := isSubsequence(s, t)
	fmt.Println("res :", res)
}
